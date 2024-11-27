package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	logginginterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/logging"
	recoveryinterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/recovery"
	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/server"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/filter"
	"github.com/AlparslanKaraguney/trux-task/pkg/grpcserver"
	"github.com/AlparslanKaraguney/trux-task/pkg/logging/logger"
	pb "github.com/AlparslanKaraguney/trux-task/proto"
)

func main() {
	// Initialize logger
	log := logger.GetLogger()

	// Validate required environment variables
	validateEnvVars([]string{
		"DATABASE_HOST", "DATABASE_USER", "DATABASE_PASS",
		"DATABASE_NAME", "DATABASE_PORT", "GRPC_SERVER_PORT",
		"HTTP_SERVER_PORT", "ENV", "DATABASE_LOG_LEVEL",
	})

	// Initialize database connection
	connection := storage.Connection(log)
	store, err := storage.NewStorage(connection)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer closeDatabaseConnection(connection, log)

	// Start gRPC server
	grpcServer, grpcAddr := setupGRPCServer(log, store)
	go startGRPCServer(grpcServer, grpcAddr, log)

	// Start HTTP health check server
	httpAddr := ":" + getEnv("HTTP_SERVER_PORT", "8080")
	go startHTTPServer(httpAddr, log)

	// Generate sample data
	generateSampleData(store)

	// Handle graceful shutdown
	handleGracefulShutdown(grpcServer, connection, log, func(code int) {
		os.Exit(code)
	})
}

func setupGRPCServer(log *logrus.Logger, store storage.Storage) (grpcserver.GRPCServer, string) {
	grpcServer := &grpcserver.RealGRPCServer{Server: grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logginginterceptor.LoggingInterceptor(log),
			recovery.UnaryServerInterceptor(recoveryinterceptor.Opts...),
		),
	)}

	svcServer := &server.SmartServiceServer{
		Storage: store,
	}
	pb.RegisterSmartServiceServer(grpcServer.Server, svcServer)
	// Enable server reflection
	reflection.Register(grpcServer.Server)

	grpcAddr := ":" + getEnv("GRPC_SERVER_PORT", "50051")

	return grpcServer.Server, grpcAddr
}

func startGRPCServer(grpcServer grpcserver.GRPCServer, addr string, log *logrus.Logger) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	log.Infof("gRPC server is listening on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func startHTTPServer(addr string, log *logrus.Logger) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Infof("HTTP health check server running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// handleGracefulShutdown handles clean shutdown on SIGINT or SIGTERM
func handleGracefulShutdown(grpcServer grpcserver.GRPCServer, connection *gorm.DB, logger *logrus.Logger, exitFunc func(code int)) {
	// Create a channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-sigChan
	logger.Infof("Received terminate signal: %v, starting graceful shutdown...", sig)

	// Gracefully stop the gRPC server
	grpcServer.GracefulStop()
	logger.Info("gRPC server stopped.")

	// Close database connection
	closeDatabaseConnection(connection, logger)

	logger.Info("Shutdown complete.")
	exitFunc(0) // Use callback instead of os.Exit to allow testing
}

// closeDatabaseConnection closes the database connection gracefully
func closeDatabaseConnection(connection *gorm.DB, logger *logrus.Logger) {
	sqlDB, err := connection.DB()
	if err != nil {
		logger.Errorf("Error retrieving database connection: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		logger.Errorf("Error closing database connection: %v", err)
	} else {
		logger.Info("Database connection closed successfully.")
	}
}

func validateEnvVars(vars []string) {
	for _, v := range vars {
		if os.Getenv(v) == "" {
			log.Panicf("Environment variable %s is required but not set.", v)
		}
	}
}

func generateSampleData(store storage.Storage) {
	// check if the data already exists
	modelList, _, err := store.ListSmartModels(&filter.SmartModelFilter{
		Limit: 1,
	})
	if err != nil {
		log.Fatalf("Failed to list smart models: %v", err)
	}

	if len(modelList) > 0 {
		log.Println("Sample data already exists. Skipping data generation.")
		return
	}

	model := &models.SmartModel{
		Name:       "Smart Engine Monitor",
		Identifier: "engine-monitor-001",
		Type:       "Device",
		Category:   "Engine",
	}
	store.CreateSmartModel(model)

	features := []*models.SmartFeature{
		{
			Name:          "Real-Time RPM Monitor",
			Identifier:    "rpm-monitor-001",
			Functionality: "Monitors engine RPM in real-time",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Real-Time Fuel Consumption Monitor",
			Identifier:    "fuel-monitor-001",
			Functionality: "Monitors fuel consumption in real-time",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Real-Time Engine Temperature Monitor",
			Identifier:    "temp-monitor-001",
			Functionality: "Monitors engine temperature in real-time",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Battery Health Monitoring",
			Identifier:    "battery-health-001",
			Functionality: "Monitors the health of the car's battery and alerts the driver to potential issues",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model
	model = &models.SmartModel{
		Name:       "Smart Tire Pressure Monitor",
		Identifier: "tire-monitor-001",
		Type:       "Device",
		Category:   "Tire",
	}
	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Real-Time Tire Pressure Monitor",
			Identifier:    "tire-pressure-monitor-001",
			Functionality: "Monitors tire pressure in real-time",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Real-Time Tire Temperature Monitor",
			Identifier:    "tire-temp-monitor-001",
			Functionality: "Monitors tire temperature in real-time",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Real-Time Tire Wear Monitor",
			Identifier:    "tire-wear-monitor-001",
			Functionality: "Monitors tire wear in real-time",
			SmartModelID:  model.ID,
		},
	}

	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Smart Infotainment System",
		Identifier: "infotainment-001",
		Type:       "Device",
		Category:   "Entertainment",
	}
	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Music Streaming",
			Identifier:    "music-streaming-001",
			Functionality: "Streams music from online services",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Navigation System",
			Identifier:    "navigation-system-001",
			Functionality: "Provides real-time GPS navigation",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Voice Command",
			Identifier:    "voice-command-001",
			Functionality: "Enables control of the system through voice commands",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Bluetooth Connectivity",
			Identifier:    "bluetooth-001",
			Functionality: "Allows pairing with smartphones for hands-free calling and media streaming",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Navigation",
			Identifier:    "navigation-001",
			Functionality: "Provides GPS navigation and real-time traffic updates",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Smart Climate Control",
		Identifier: "climate-control-001",
		Type:       "Device",
		Category:   "Comfort",
	}
	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Temperature Control",
			Identifier:    "temp-control-001",
			Functionality: "Adjusts the cabin temperature automatically",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Air Quality Monitor",
			Identifier:    "air-quality-001",
			Functionality: "Tracks cabin air quality and filters pollutants",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Remotely Controllable Camera",
		Identifier: "remote-camera-001",
		Type:       "Device",
		Category:   "Surveillance",
	}
	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Take Screenshot",
			Identifier:    "camera-screenshot-001",
			Functionality: "Captures a still image from the camera",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Get Live Video URL",
			Identifier:    "live-video-url-001",
			Functionality: "Retrieves the live video stream URL",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Pan and Tilt Control",
			Identifier:    "pan-tilt-control-001",
			Functionality: "Allows remote control of the camera's pan and tilt",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Open Weather Map",
		Identifier: "weather-service-001",
		Type:       "Service",
		Category:   "Weather",
	}

	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Get Weekly Forecast",
			Identifier:    "weekly-forecast-001",
			Functionality: "Retrieves the weather forecast for the next 7 days",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Get Daily Forecast",
			Identifier:    "daily-forecast-001",
			Functionality: "Retrieves the weather forecast for a specific day",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Get Current Weather",
			Identifier:    "current-weather-001",
			Functionality: "Provides current weather conditions for a city",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "IMDB Movie Database",
		Identifier: "imdb-service-001",
		Type:       "Service",
		Category:   "Entertainment",
	}

	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Search Movie",
			Identifier:    "search-movie-001",
			Functionality: "Searches for a movie by title",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Get Top Rated Movies",
			Identifier:    "top-rated-001",
			Functionality: "Retrieves a list of top-rated movies",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Get Movie Details",
			Identifier:    "movie-details-001",
			Functionality: "Provides detailed information about a movie",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Advanced Driver Assistance System",
		Identifier: "adas-001",
		Type:       "Device",
		Category:   "Safety",
	}

	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Lane Keep Assist",
			Identifier:    "lane-keep-001",
			Functionality: "Detects lane markers and adjusts steering to keep the car in its lane",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Adaptive Cruise Control",
			Identifier:    "adaptive-cruise-001",
			Functionality: "Automatically adjusts the car's speed to maintain a safe following distance",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Automatic Emergency Braking",
			Identifier:    "auto-brake-001",
			Functionality: "Applies brakes automatically if an imminent collision is detected",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

	// Create another model

	model = &models.SmartModel{
		Name:       "Parking Assistance System",
		Identifier: "parking-assist-001",
		Type:       "Device",
		Category:   "Convenience",
	}

	store.CreateSmartModel(model)

	features = []*models.SmartFeature{
		{
			Name:          "Reverse Camera",
			Identifier:    "reverse-camera-001",
			Functionality: "Provides a rear-view video feed to assist with parking",
			SmartModelID:  model.ID,
		},
		{
			Name:          "Automatic Parking",
			Identifier:    "auto-parking-001",
			Functionality: "Guides the car into a parking spot with minimal driver input",
			SmartModelID:  model.ID,
		},
	}
	for _, f := range features {
		store.CreateSmartFeature(f)
	}

}
