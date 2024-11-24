package server

import (
	"context"
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	pb "github.com/AlparslanKaraguney/trux-task/proto"
	"github.com/jackc/pgx/v5/pgconn"

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
		// Handle PostgreSQL duplicate key error since we are using PostgreSQL and gorm natively does not support this error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "SmartModel already exists: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to create SmartModel: %v", err)
	}
	return &pb.SmartModelResponse{Model: convertToProtoSmartModel(model)}, nil
}

func (s *SmartServiceServer) GetSmartModel(ctx context.Context, req *pb.SmartModelQuery) (*pb.SmartModelResponse, error) {

	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	model, err := s.Storage.GetSmartModel(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartModel: %v", err)
	}
	return &pb.SmartModelResponse{Model: convertToProtoSmartModel(model)}, nil
}

// Update SmartModel
func (s *SmartServiceServer) UpdateSmartModel(ctx context.Context, req *pb.SmartModelRequest) (*pb.SmartModelResponse, error) {
	model := req.Model

	// Validate request
	if model.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	// Update the model in storage
	err := s.Storage.UpdateSmartModel(&models.SmartModel{
		ID:         model.Id,
		Name:       model.Name,
		Identifier: model.Identifier,
		Type:       model.Type,
		Category:   model.Category,
	})
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		// Handle PostgreSQL duplicate key error since we are using PostgreSQL and gorm natively does not support this error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "SmartModel already exists: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "Failed to update SmartModel: %v", err)
	}

	return &pb.SmartModelResponse{Model: model}, nil
}

// Delete SmartModel
func (s *SmartServiceServer) DeleteSmartModel(ctx context.Context, req *pb.SmartModelQuery) (*pb.DeleteResponse, error) {
	// Validate request
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	// Delete the model in storage
	err := s.Storage.DeleteSmartModel(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete SmartModel: %v", err)
	}

	return &pb.DeleteResponse{Message: "SmartModel deleted successfully", Success: true}, nil
}

func (s *SmartServiceServer) CreateSmartFeature(ctx context.Context, req *pb.SmartFeatureRequest) (*pb.SmartFeatureResponse, error) {
	feature := &models.SmartFeature{
		Name:          req.Feature.Name,
		Identifier:    req.Feature.Identifier,
		Functionality: req.Feature.Functionality,
		SmartModelID:  req.Feature.SmartModelId,
	}
	err := s.Storage.CreateSmartFeature(feature)
	if err != nil {
		// Handle PostgreSQL duplicate key error since we are using PostgreSQL and gorm natively does not support this error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "SmartFeature already exists: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "Failed to create SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}

func (s *SmartServiceServer) GetSmartFeature(ctx context.Context, req *pb.SmartFeatureQuery) (*pb.SmartFeatureResponse, error) {
	// Validate request
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	feature, err := s.Storage.GetSmartFeature(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}

// Update SmartFeature
func (s *SmartServiceServer) UpdateSmartFeature(ctx context.Context, req *pb.SmartFeatureRequest) (*pb.SmartFeatureResponse, error) {
	feature := req.Feature

	// Validate request
	if feature.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	// Update the feature in storage
	err := s.Storage.UpdateSmartFeature(&models.SmartFeature{
		ID:            feature.Id,
		Name:          feature.Name,
		Identifier:    feature.Identifier,
		Functionality: feature.Functionality,
		SmartModelID:  feature.SmartModelId,
	})
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		// Handle PostgreSQL duplicate key error since we are using PostgreSQL and gorm natively does not support this error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "SmartFeature already exists: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "Failed to update SmartFeature: %v", err)
	}

	return &pb.SmartFeatureResponse{Feature: feature}, nil
}

// Delete SmartFeature
func (s *SmartServiceServer) DeleteSmartFeature(ctx context.Context, req *pb.SmartFeatureQuery) (*pb.DeleteResponse, error) {
	// Validate request
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	// Delete the feature in storage
	err := s.Storage.DeleteSmartFeature(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete SmartFeature: %v", err)
	}

	return &pb.DeleteResponse{Message: "SmartFeature deleted successfully", Success: true}, nil
}
