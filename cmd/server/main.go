package main

import (
	"log"
	"net"
	"os"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/server"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	pb "github.com/AlparslanKaraguney/trux-task/proto"

	loginterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/logging"
	recoveryinterceptor "github.com/AlparslanKaraguney/trux-task/interceptors/recovery"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Redirect standard log to logrus
	log.SetOutput(logger.Writer())
	log.SetFlags(0) // Disable default timestamp since logrus adds it

	connection := storage.Connection(logger)

	// Initialize storage
	store, err := storage.NewStorage(connection)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			loginterceptor.LoggingInterceptor(logger),
			recovery.UnaryServerInterceptor(recoveryinterceptor.Opts...),
		),
	)

	svcServer := &server.SmartServiceServer{
		Storage: store,
	}
	pb.RegisterSmartServiceServer(grpcServer, svcServer)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// Sample data generation
	generateSampleData(store)
}

func generateSampleData(store storage.Storage) {
	model := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	store.CreateSmartModel(model)

	feature := &models.SmartFeature{
		Name:          "Get Heart Rate",
		Identifier:    "sw-hr-001",
		Functionality: "Retrieve current heart rate",
		ModelID:       model.ID,
	}
	store.CreateSmartFeature(feature)
}
