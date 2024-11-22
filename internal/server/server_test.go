package server

import (
	"context"
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/mocks"
	pb "github.com/AlparslanKaraguney/trux-task/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sminstance := &models.SmartModel{
		ID:         1,
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
		Features: []models.SmartFeature{
			{
				ID:            1,
				Name:          "Future",
				Identifier:    "f-001",
				Functionality: "Predict the future",
				SmartModelID:  1,
			},
		},
	}

	mockStorage.
		EXPECT().
		GetSmartModel(sminstance.ID).
		Return(sminstance, nil).
		Times(1)

	server := &SmartServiceServer{
		Storage: mockStorage,
	}

	ctx := context.Background()

	req := &pb.SmartModelQuery{Id: sminstance.ID}

	res, err := server.GetSmartModel(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, res.Model.Id)
}

func TestCreateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	// Define the behavior of CreateSmartModel
	mockStorage.
		EXPECT().
		CreateSmartModel(sminstance).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{
		Storage: mockStorage,
	}

	ctx := context.Background()

	// Create a SmartModel request
	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Name:       sminstance.Name,
			Identifier: sminstance.Identifier,
			Type:       sminstance.Type,
			Category:   sminstance.Category,
		},
	}

	// Call CreateSmartModel
	res, err := server.CreateSmartModel(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Model.Identifier, res.Model.Identifier)
}

func TestUpdateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sminstance := &models.SmartModel{
		ID:         1,
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	mockStorage.
		EXPECT().
		UpdateSmartModel(sminstance).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Id:         sminstance.ID,
			Name:       sminstance.Name,
			Identifier: sminstance.Identifier,
			Type:       sminstance.Type,
			Category:   sminstance.Category,
		},
	}

	res, err := server.UpdateSmartModel(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Model.Name, res.Model.Name)

}

func TestDeleteSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sminstance := &models.SmartModel{
		ID: 1,
	}

	mockStorage.
		EXPECT().
		DeleteSmartModel(sminstance.ID).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelQuery{Id: sminstance.ID}

	res, err := server.DeleteSmartModel(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sfinstance := &models.SmartFeature{
		ID: 1,
	}

	mockStorage.
		EXPECT().
		GetSmartFeature(sfinstance.ID).
		Return(sfinstance, nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: sfinstance.ID}

	res, err := server.GetSmartFeature(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, res.Feature.Id)
}

func TestCreateSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sfinstance := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  1,
	}

	mockStorage.
		EXPECT().
		CreateSmartFeature(sfinstance).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Name:          sfinstance.Name,
			Identifier:    sfinstance.Identifier,
			Functionality: sfinstance.Functionality,
			SmartModelId:  sfinstance.SmartModelID,
		},
	}

	res, err := server.CreateSmartFeature(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Feature.Name, res.Feature.Name)
}

func TestUpdateSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sfinstance := &models.SmartFeature{
		ID:            1,
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  1,
	}

	mockStorage.
		EXPECT().
		UpdateSmartFeature(sfinstance).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Id:            sfinstance.ID,
			Name:          sfinstance.Name,
			Identifier:    sfinstance.Identifier,
			Functionality: sfinstance.Functionality,
			SmartModelId:  sfinstance.SmartModelID,
		},
	}

	res, err := server.UpdateSmartFeature(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Feature.Name, res.Feature.Name)
}

func TestDeleteSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	sfinstance := &models.SmartFeature{
		ID: 1,
	}

	mockStorage.
		EXPECT().
		DeleteSmartFeature(sfinstance.ID).
		Return(nil).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: sfinstance.ID}

	res, err := server.DeleteSmartFeature(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
