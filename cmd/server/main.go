package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	logginginterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/logging"
	recoveryinterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/recovery"
	"github.com/AlparslanKaraguney/trux-task/internal/server"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
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

	fmt.Println("addr", addr)

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

// func generateSampleData(store storage.Storage) {
// 	model := &models.SmartModel{
// 		Name:       "Smart Watch",
// 		Identifier: "sw-001",
// 		Type:       "Device",
// 		Category:   "Wearable",
// 	}
// 	store.CreateSmartModel(model)

// 	feature := &models.SmartFeature{
// 		Name:          "Get Heart Rate",
// 		Identifier:    "sw-hr-001",
// 		Functionality: "Retrieve current heart rate",
// 		SmartModelID:  model.ID,
// 	}
// 	store.CreateSmartFeature(feature)
// }
