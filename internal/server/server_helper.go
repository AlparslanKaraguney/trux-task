package server

import (
	"github.com/AlparslanKaraguney/trux-task/internal/models"
	pb "github.com/AlparslanKaraguney/trux-task/proto"
)

func convertToProtoSmartFeature(feature *models.SmartFeature) *pb.SmartFeature {
	return &pb.SmartFeature{
		Id:            feature.ID,
		Name:          feature.Name,
		Identifier:    feature.Identifier,
		Functionality: feature.Functionality,
		ModelId:       feature.ModelID,
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
