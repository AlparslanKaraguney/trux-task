package server

import (
	"context"
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/mocks"
	pb "github.com/AlparslanKaraguney/trux-task/proto"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func TestCreateDublicateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the PostgreSQL error (duplicate key violation)
	pgError := &pgconn.PgError{
		Code:    "23505", // Unique constraint violation code
		Message: "duplicate key value violates unique constraint",
	}

	mockStorage.
		EXPECT().
		CreateSmartModel(gomock.Any()).
		Return(pgError).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	_, err := server.CreateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is AlreadyExists gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.AlreadyExists, status.Code())
	}
}

func TestUpdateSmartModelNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		UpdateSmartModel(gomock.Any()).
		Return(storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Id:         1,
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	_, err := server.UpdateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestUpdateDublicateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the PostgreSQL error (duplicate key violation)
	pgError := &pgconn.PgError{
		Code:    "23505", // Unique constraint violation code
		Message: "duplicate key value violates unique constraint",
	}

	mockStorage.
		EXPECT().
		UpdateSmartModel(gomock.Any()).
		Return(pgError).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Id:         1,
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	_, err := server.UpdateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is AlreadyExists gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.AlreadyExists, status.Code())
	}
}

func TestDeleteSmartModelNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		DeleteSmartModel(gomock.Any()).
		Return(storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelQuery{Id: 1}

	_, err := server.DeleteSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestDeleteSmartFeatureNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		DeleteSmartFeature(gomock.Any()).
		Return(storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: 1}

	_, err := server.DeleteSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestCreateDublicateSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the PostgreSQL error (duplicate key violation)
	pgError := &pgconn.PgError{
		Code:    "23505", // Unique constraint violation code
		Message: "duplicate key value violates unique constraint",
	}

	mockStorage.
		EXPECT().
		CreateSmartFeature(gomock.Any()).
		Return(pgError).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.CreateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is AlreadyExists gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.AlreadyExists, status.Code())
	}
}

func TestUpdateSmartFeatureNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		UpdateSmartFeature(gomock.Any()).
		Return(storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Id:            1,
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.UpdateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestGetSmartModelNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		GetSmartModel(gomock.Any()).
		Return(nil, storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelQuery{Id: 1}

	_, err := server.GetSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestUpdateSmartFeature_AlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the PostgreSQL error (duplicate key violation)
	pgError := &pgconn.PgError{
		Code:    "23505", // Unique constraint violation code
		Message: "duplicate key value violates unique constraint",
	}

	mockStorage.
		EXPECT().
		UpdateSmartFeature(gomock.Any()).
		Return(pgError).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Id:            1,
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.UpdateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is AlreadyExists gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.AlreadyExists, status.Code())
	}
}

func TestCreateSmartFeature_AlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the PostgreSQL error (duplicate key violation)
	pgError := &pgconn.PgError{
		Code:    "23505", // Unique constraint violation code
		Message: "duplicate key value violates unique constraint",
	}

	mockStorage.
		EXPECT().
		CreateSmartFeature(gomock.Any()).
		Return(pgError).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.CreateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is AlreadyExists gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.AlreadyExists, status.Code())
	}
}

func TestCreateSmartModel_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		CreateSmartModel(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	_, err := server.CreateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestGetSmartModel_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		GetSmartModel(gomock.Any()).
		Return(nil, status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelQuery{Id: 1}

	_, err := server.GetSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestUpdateSmartModel_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		UpdateSmartModel(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{
			Id:         1,
			Name:       "Smart Watch",
			Identifier: "sw-001",
			Type:       "Device",
			Category:   "Wearable",
		},
	}

	_, err := server.UpdateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestGetSmartModel_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartModelQuery{}

	_, err := server.GetSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestUpdateSmartModel_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartModelRequest{
		Model: &pb.SmartModel{},
	}

	_, err := server.UpdateSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestDeleteSmartModel_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartModelQuery{}

	_, err := server.DeleteSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestGetSmartFeature_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartFeatureQuery{}

	_, err := server.GetSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestCreateSmartFeature_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		CreateSmartFeature(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.CreateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestUpdateSmartFeature_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{},
	}

	_, err := server.UpdateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestUpdateSmartFeature_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		UpdateSmartFeature(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureRequest{
		Feature: &pb.SmartFeature{
			Id:            1,
			Name:          "Future",
			Identifier:    "f-001",
			Functionality: "Predict the future",
			SmartModelId:  1,
		},
	}

	_, err := server.UpdateSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestDeleteSmartFeature_IdRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := &SmartServiceServer{}

	req := &pb.SmartFeatureQuery{}

	_, err := server.DeleteSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is InvalidArgument gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.InvalidArgument, status.Code())
	}
}

func TestDeleteSmartModel_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		DeleteSmartModel(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartModelQuery{Id: 1}

	_, err := server.DeleteSmartModel(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestGetSmartFeature_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (not found)
	mockStorage.
		EXPECT().
		GetSmartFeature(gomock.Any()).
		Return(nil, storage.ErrNotFound).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: 1}

	_, err := server.GetSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is NotFound gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.NotFound, status.Code())
	}
}

func TestGetSmartFeature_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		GetSmartFeature(gomock.Any()).
		Return(nil, status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: 1}

	_, err := server.GetSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

func TestDeleteSmartFeature_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	// Simulate the storage error (internal error)
	mockStorage.
		EXPECT().
		DeleteSmartFeature(gomock.Any()).
		Return(status.Errorf(codes.Internal, "Internal error")).
		Times(1)

	server := &SmartServiceServer{Storage: mockStorage}

	req := &pb.SmartFeatureQuery{Id: 1}

	_, err := server.DeleteSmartFeature(context.Background(), req)
	assert.Error(t, err)
	// assert the error is Internal gives the correct error code
	if status, ok := status.FromError(err); ok {
		assert.Equal(t, codes.Internal, status.Code())
	}
}

// func TestListSmartModels(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStorage := mocks.NewMockStorage(ctrl)

// 	sminstance := &models.SmartModel{
// 		ID:         1,
// 		Name:       "Smart Watch",
// 		Identifier: "sw-001",
// 		Type:       "Device",
// 		Category:   "Wearable",
// 		Features: []models.SmartFeature{
// 			{
// 				ID:            1,
// 				Name:          "Future",
// 				Identifier:    "f-001",
// 				Functionality: "Predict the future",
// 				SmartModelID:  1,
// 			},
// 		},
// 	}

// 	mockStorage.
// 		EXPECT().
// 		ListSmartModels().
// 		Return([]*models.SmartModel{sminstance}, nil).
// 		Times(1)

// 	server := &SmartServiceServer{
// 		Storage: mockStorage,
// 	}

// 	ctx := context.Background()

// 	req := &pb.Empty{}

// 	res, err := server.ListSmartModels(ctx, req)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, len(res.Models))
// }

// func TestListSmartFeatures(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStorage := mocks.NewMockStorage(ctrl)

// 	sfinstance := &models.SmartFeature{
// 		ID:            1,
// 		Name:          "Future",
// 		Identifier:    "f-001",
// 		Functionality: "Predict the future",
// 		SmartModelID:  1,
// 	}

// 	mockStorage.
// 		EXPECT().
// 		ListSmartFeatures().
// 		Return([]*models.SmartFeature{sfinstance}, nil).
// 		Times(1)

// 	server := &SmartServiceServer{
// 		Storage: mockStorage,
// 	}

// 	ctx := context.Background()

// 	req := &pb.Empty{}

// 	res, err := server.ListSmartFeatures(ctx, req)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, len(res.Features))
// }
