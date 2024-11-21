package server

import (
	"context"
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/mocks"
	"github.com/AlparslanKaraguney/trux-task/internal/models"
	pb "github.com/AlparslanKaraguney/trux-task/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Define the behavior of CreateSmartModel
	mockStorage.
		EXPECT().
		CreateSmartModel(gomock.Any()).
		Return(nil).
		Times(1)

	// Define the behavior of GetSmartModel
	mockStorage.
		EXPECT().
		GetSmartModel("sw-001").
		Return(&models.SmartModel{
			ID:         1,
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
			Features: []models.SmartFeature{
				{
					ID:            1,
					Name:          "Get Heart Rate",
					Identifier:    "sw-hr-001",
					Functionality: "Retrieve current heart rate",
					ModelID:       1,
				},
				{
					ID:            2,
					Name:          "Get Heart Rate 2",
					Identifier:    "sw-hr-002",
					Functionality: "Retrieve current heart rate 2",
					ModelID:       1,
				},
			},
		}, nil).
		Times(1)

	svcServer := &SmartServiceServer{
		Storage: mockStorage,
	}

	ctx := context.Background()

	// Create a SmartModel request
	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	// Call CreateSmartModel
	res, err := svcServer.CreateSmartModel(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Model.Name, res.Model.Name)

	// Call GetSmartModel
	query := &pb.SmartModelQuery{Identifier: "sw-001"}
	getRes, err := svcServer.GetSmartModel(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, req.Model.Name, getRes.Model.Name)
}
