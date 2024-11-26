package server

import (
	"github.com/AlparslanKaraguney/trux-task/internal/entities"
	"github.com/AlparslanKaraguney/trux-task/internal/models"
	pb "github.com/AlparslanKaraguney/trux-task/proto"
)

func convertToProtoSmartFeature(feature *models.SmartFeature) *pb.SmartFeature {
	return &pb.SmartFeature{
		Id:            feature.ID,
		Name:          feature.Name,
		Identifier:    feature.Identifier,
		Functionality: feature.Functionality,
		SmartModelId:  feature.SmartModelID,
	}
}

func convertToProtoSmartModel(model *models.SmartModel) *pb.SmartModel {
	features := make([]*pb.SmartFeature, len(model.Features))
	for i, f := range model.Features {
		features[i] = convertToProtoSmartFeature(&f)
	}
	return &pb.SmartModel{
		Id:         model.ID,
		Name:       model.Name,
		Identifier: model.Identifier,
		Type:       model.Type,
		Category:   model.Category,
		Features:   features,
	}
}

func convertToProtoSmartModels(models []models.SmartModel) []*pb.SmartModel {
	var protoModels []*pb.SmartModel
	for _, model := range models {
		protoModels = append(protoModels, &pb.SmartModel{
			Id:         model.ID,
			Name:       model.Name,
			Identifier: model.Identifier,
			Type:       model.Type,
			Category:   model.Category,
		})
	}
	return protoModels
}

func convertToProtoSmartFeatures(features []models.SmartFeature) []*pb.SmartFeature {
	var protoFeatures []*pb.SmartFeature
	for _, feature := range features {
		protoFeatures = append(protoFeatures, &pb.SmartFeature{
			Id:            feature.ID,
			Name:          feature.Name,
			Identifier:    feature.Identifier,
			Functionality: feature.Functionality,
			SmartModelId:  feature.SmartModelID,
		})
	}
	return protoFeatures
}

func convertToProtoPagination(pagination *entities.Pagination) *pb.Pagination {
	return &pb.Pagination{
		TotalRows: int32(pagination.TotalRows),
		Limit:     int32(pagination.Limit),
		Offset:    int32(pagination.Offset),
	}
}
