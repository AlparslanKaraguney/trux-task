package server

import (
	"context"
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	pb "github.com/AlparslanKaraguney/trux-task/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SmartServiceServer struct {
	pb.UnimplementedSmartServiceServer
	Storage storage.Storage
}

func (s *SmartServiceServer) CreateSmartModel(ctx context.Context, req *pb.SmartModelRequest) (*pb.SmartModelResponse, error) {
	model := &models.SmartModel{
		Name:       req.Model.Name,
		Identifier: req.Model.Identifier,
		Type:       req.Model.Type,
		Category:   req.Model.Category,
	}
	err := s.Storage.CreateSmartModel(model)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create SmartModel: %v", err)
	}
	return &pb.SmartModelResponse{Model: convertToProtoSmartModel(model)}, nil
}

func (s *SmartServiceServer) GetSmartModel(ctx context.Context, req *pb.SmartModelQuery) (*pb.SmartModelResponse, error) {
	model, err := s.Storage.GetSmartModel(req.Identifier)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartModel: %v", err)
	}
	return &pb.SmartModelResponse{Model: convertToProtoSmartModel(model)}, nil
}

func (s *SmartServiceServer) CreateSmartFeature(ctx context.Context, req *pb.SmartFeatureRequest) (*pb.SmartFeatureResponse, error) {
	feature := &models.SmartFeature{
		Name:          req.Feature.Name,
		Identifier:    req.Feature.Identifier,
		Functionality: req.Feature.Functionality,
		ModelID:       req.Feature.ModelId,
	}
	err := s.Storage.CreateSmartFeature(feature)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}

func (s *SmartServiceServer) GetSmartFeature(ctx context.Context, req *pb.SmartFeatureQuery) (*pb.SmartFeatureResponse, error) {
	feature, err := s.Storage.GetSmartFeature(req.Identifier)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}
