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

	// smartFuture := &pb.SmartFeature{
	// 	Name:          "Future",
	// 	Identifier:    "f-001",
	// 	Functionality: "Predict the future",
	// 	SmartModelId:  1,
	// }

	// res, err := client.CreateSmartFeature(ctx, &pb.SmartFeatureRequest{Feature: smartFuture})
	// if err != nil {
	// 	log.Fatalf("Could not create SmartFeature: %v", err)
	// }

	// log.Printf("Created SmartFeature: %v", res.Feature)

	// Create a SmartModel
	// modelReq := &pb.SmartModelRequest{
	// 	Model: &pb.SmartModel{
	// 		Name:       "Smart Watch 1",
	// 		Identifier: "sw-002",
	// 		Type:       "Device",
	// 		Category:   "Wearable",
	// 	},
	// }

	// createModelRes, err := client.CreateSmartModel(ctx, modelReq)
	// if err != nil {
	// 	log.Fatalf("Could not create SmartModel: %v", err)
	// }

	// Retrieve the SmartModel
	modelQuery := &pb.SmartModelQuery{Id: 1}
	getModelRes, err := client.GetSmartModel(ctx, modelQuery)
	if err != nil {
		log.Fatalf("Could not get SmartModel: %v", err)
	}

	log.Println(getModelRes.Model.Features)
	log.Printf("Retrieved SmartModel: %v", getModelRes.Model)
}
