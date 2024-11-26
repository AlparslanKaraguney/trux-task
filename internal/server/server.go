package server

import (
	"context"
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/filter"
	apperrors "github.com/AlparslanKaraguney/trux-task/pkg/errors"
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
		if errors.Is(err, apperrors.ErrAlreadyExists) {
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
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartModel: %v", err)
	}
	return &pb.SmartModelResponse{Model: convertToProtoSmartModel(model)}, nil
}

func (s *SmartServiceServer) UpdateSmartModel(ctx context.Context, req *pb.SmartModelRequest) (*pb.SmartModelResponse, error) {
	model := req.Model

	if model.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	err := s.Storage.UpdateSmartModel(&models.SmartModel{
		ID:         model.Id,
		Name:       model.Name,
		Identifier: model.Identifier,
		Type:       model.Type,
		Category:   model.Category,
	})

	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		if errors.Is(err, apperrors.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "SmartModel already exists: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to update SmartModel: %v", err)
	}

	return &pb.SmartModelResponse{Model: model}, nil
}

func (s *SmartServiceServer) DeleteSmartModel(ctx context.Context, req *pb.SmartModelQuery) (*pb.DeleteResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	err := s.Storage.DeleteSmartModel(req.Id)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartModel not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete SmartModel: %v", err)
	}

	return &pb.DeleteResponse{Message: "SmartModel deleted successfully", Success: true}, nil
}

func (s *SmartServiceServer) ListSmartModel(ctx context.Context, req *pb.SmartModelListQuery) (*pb.SmartModelListResponse, error) {

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	filter := &filter.SmartModelFilter{
		Identifier: req.Identifier,
		Name:       req.Name,
		Type:       req.Type,
		Category:   req.Category,
		Limit:      req.Limit,
		Offset:     req.Offset,
		OrderBy:    req.OrderBy,
	}

	models, pagination, err := s.Storage.ListSmartModels(filter)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidOrderBy) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid OrderBy argument: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to list SmartModels: %v", err)
	}

	return &pb.SmartModelListResponse{
		Data:       convertToProtoSmartModels(models),
		Pagination: convertToProtoPagination(pagination),
	}, nil

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
		if errors.Is(err, apperrors.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "SmartModel already exists: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "Failed to create SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}

func (s *SmartServiceServer) GetSmartFeature(ctx context.Context, req *pb.SmartFeatureQuery) (*pb.SmartFeatureResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	feature, err := s.Storage.GetSmartFeature(req.Id)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get SmartFeature: %v", err)
	}
	return &pb.SmartFeatureResponse{Feature: convertToProtoSmartFeature(feature)}, nil
}

func (s *SmartServiceServer) UpdateSmartFeature(ctx context.Context, req *pb.SmartFeatureRequest) (*pb.SmartFeatureResponse, error) {
	feature := req.Feature

	if feature.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	err := s.Storage.UpdateSmartFeature(&models.SmartFeature{
		ID:            feature.Id,
		Name:          feature.Name,
		Identifier:    feature.Identifier,
		Functionality: feature.Functionality,
		SmartModelID:  feature.SmartModelId,
	})
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		if errors.Is(err, apperrors.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "SmartModel already exists: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to update SmartFeature: %v", err)
	}

	return &pb.SmartFeatureResponse{Feature: feature}, nil
}

func (s *SmartServiceServer) DeleteSmartFeature(ctx context.Context, req *pb.SmartFeatureQuery) (*pb.DeleteResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	err := s.Storage.DeleteSmartFeature(req.Id)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "SmartFeature not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete SmartFeature: %v", err)
	}

	return &pb.DeleteResponse{Message: "SmartFeature deleted successfully", Success: true}, nil
}

func (s *SmartServiceServer) ListSmartFeature(ctx context.Context, req *pb.SmartFeatureListQuery) (*pb.SmartFeatureListResponse, error) {

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	filter := &filter.SmartFeatureFilter{
		Identifier:    req.Identifier,
		Name:          req.Name,
		Functionality: req.Functionality,
		SmartModelId:  req.SmartModelId,
		Limit:         req.Limit,
		Offset:        req.Offset,
		OrderBy:       req.OrderBy,
	}

	features, pagination, err := s.Storage.ListSmartFeatures(filter)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidOrderBy) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid OrderBy argument: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to list SmartFeatures: %v", err)
	}

	return &pb.SmartFeatureListResponse{
		Data:       convertToProtoSmartFeatures(features),
		Pagination: convertToProtoPagination(pagination),
	}, nil

}
