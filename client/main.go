package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AlparslanKaraguney/trux-task/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSmartServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// // Create a SmartFeature
	// featureReq := &pb.SmartFeatureRequest{
	// 	Feature: &pb.SmartFeature{
	// 		Name:          "Get Heart Rate",
	// 		Identifier:    "sw-hr-001",
	// 		Functionality: "Retrieve current heart rate",
	// 		ModelId:       1,
	// 	},
	// }

	// featureRes, err := client.CreateSmartFeature(ctx, featureReq)
	// if err != nil {
	// 	log.Fatalf("Could not create SmartFeature: %v", err)
	// }

	// log.Printf("Created SmartFeature: %v", featureRes.Feature)

	// // Create a SmartFeature
	// featureReq2 := &pb.SmartFeatureRequest{
	// 	Feature: &pb.SmartFeature{
	// 		Name:          "Get Heart Rate 2",
	// 		Identifier:    "sw-hr-002",
	// 		Functionality: "Retrieve current heart rate 2",
	// 		ModelId:       1,
	// 	},
	// }

	// featureRes2, err := client.CreateSmartFeature(ctx, featureReq2)
	// if err != nil {
	// 	log.Fatalf("Could not create SmartFeature: %v", err)
	// }

	// log.Printf("Created SmartFeature: %v", featureRes2.Feature)

	// Create a SmartModel
	// modelReq := &pb.SmartModelRequest{
	// 	Model: &pb.SmartModel{
	// 		Name:       "Smart Watch 2",
	// 		Identifier: "sw-002",
	// 		Type:       "Device",
	// 		Category:   "Wearable",
	// 	},
	// }

	// modelRes, err := client.CreateSmartModel(ctx, modelReq)
	// if err != nil {
	// 	log.Fatalf("Could not create SmartModel: %v", err)
	// }
	// log.Printf("Created SmartModel: %v", modelRes.Model)

	// Retrieve the SmartModel
	modelQuery := &pb.SmartModelQuery{Identifier: "sw-002"}
	getModelRes, err := client.GetSmartModel(ctx, modelQuery)
	if err != nil {
		log.Fatalf("Could not get SmartModel: %v", err)
	}
	log.Printf("Retrieved SmartModel: %v", getModelRes.Model)
}
