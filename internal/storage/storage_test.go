package storage

import (
	"testing"
	"time"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/filter"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStorageCreateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)
}

func TestStorageGetSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)
}

func TestStorageUpdateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Update the model
	model.Name = "Smart Watch 2"
	err = mokcedStorage.UpdateSmartModel(model)
	assert.NoError(t, err)

	// Check if the model was updated
	updatedModel, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, model.Name, updatedModel.Name)
}

func TestStorageDeleteSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Delete the model
	err = mokcedStorage.DeleteSmartModel(sminstance.ID)
	assert.NoError(t, err)

	// Check if the model was deleted
	_, err = mokcedStorage.GetSmartModel(sminstance.ID)
	assert.Error(t, err)
}

func TestStorageDeleteSmartModelAndItsFutures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)

	// Delete the model
	err = mokcedStorage.DeleteSmartModel(sminstance.ID)
	assert.NoError(t, err)

	// Check if the model was deleted
	_, err = mokcedStorage.GetSmartModel(sminstance.ID)
	assert.Error(t, err)

	// Check if the future was deleted
	_, err = mokcedStorage.GetSmartFeature(feature.ID)
	assert.Error(t, err)
}

func TestStorageCreateSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)
}

func TestStorageGetSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)
}

func TestStorageUpdateSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)

	// Update the future
	feature.Name = "Future 2"
	err = mokcedStorage.UpdateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was updated
	updatedFeature, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, updatedFeature.Name)
}

func TestStorageDeleteSmartFeature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)

	// Delete the future
	err = mokcedStorage.DeleteSmartFeature(feature.ID)
	assert.NoError(t, err)

	// Check if the future was deleted
	_, err = mokcedStorage.GetSmartFeature(feature.ID)
	assert.Error(t, err)

	// Check if the model still exists
	_, err = mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
}

func TestStorageDeleteSmartFeatureAndItsModelStillExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}

	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Check if the model was created
	model, err := mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
	assert.Equal(t, sminstance.Name, model.Name)

	// Create a future
	feature := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}

	err = mokcedStorage.CreateSmartFeature(feature)
	assert.NoError(t, err)

	// Check if the future was created
	featureModel, err := mokcedStorage.GetSmartFeature(feature.ID)
	assert.NoError(t, err)
	assert.Equal(t, feature.Name, featureModel.Name)

	// Delete the future
	err = mokcedStorage.DeleteSmartFeature(feature.ID)
	assert.NoError(t, err)

	// Check if the future was deleted
	_, err = mokcedStorage.GetSmartFeature(feature.ID)
	assert.Error(t, err)

	// Check if the model still exists
	_, err = mokcedStorage.GetSmartModel(sminstance.ID)
	assert.NoError(t, err)
}

func TestStorageSmartModelDeleteNonExistItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	err = mokcedStorage.DeleteSmartModel(0)
	assert.Error(t, err)

}

func TestStorageSmartFeatureDeleteNonExistItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	err = mokcedStorage.DeleteSmartFeature(0)
	assert.Error(t, err)
}

func TestStorageSmartModelUpdateNonExistItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	err = mokcedStorage.UpdateSmartModel(&models.SmartModel{
		ID: 456456,
	})
	assert.Error(t, err)
}

func TestStorageSmartFeatureUpdateNonExistItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	err = mokcedStorage.UpdateSmartFeature(&models.SmartFeature{
		ID: 456456,
	})
	assert.Error(t, err)
}

func TestStorageListSmartModels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	// Create models
	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
}

func TestStorageListSmartModels_FilterByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	// Create models
	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{Name: "Smart Watch", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, sminstance1.Name, models[0].Name)
}

func TestStorageListSmartModels_FilterByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	// Create models
	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{Category: "Wearable", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, sminstance1.Category, models[0].Category)
}

func TestStorageListSmartModels_FilterByType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	// Create models
	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{Type: "Device", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
}

func TestStorageListSmartModels_FilterByLimitAndOffset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{Limit: 1, Offset: 1})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, sminstance1.Name, models[0].Name)
}

func TestStorageListSmartModels_OrderByAscName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{OrderBy: "name asc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
	assert.Equal(t, sminstance2.Name, models[0].Name)
}

func TestStorageListSmartModels_OrderByDescName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{OrderBy: "name desc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
	assert.Equal(t, sminstance1.Name, models[0].Name)
}

func TestStorageListSmartModels_OrderByAscCreatedAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{OrderBy: "created_at asc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
	assert.Equal(t, sminstance1.Name, models[0].Name)
}

func TestStorageListSmartModels_OrderByDescCreatedAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	time.Sleep(1 * time.Microsecond)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// List models
	models, _, err := mokcedStorage.ListSmartModels(&filter.SmartModelFilter{OrderBy: "created_at desc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(models))
	assert.Equal(t, sminstance2.Name, models[0].Name)
}

func TestStorageListSmartModels_OrderByInvalidValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	_, _, err = mokcedStorage.ListSmartModels(&filter.SmartModelFilter{OrderBy: "invalid", Limit: 10})
	assert.Error(t, err)
}

func TestStorageListSmartFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(features))
}

func TestStorageListSmartFeatures_FilterByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{Name: "Future", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}

func TestStorageListSmartFeatures_FilterByFunctionality(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{Functionality: "Predict the future", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(features))
	assert.Equal(t, feature1.Functionality, features[0].Functionality)
}

func TestStorageListSmartFeatures_FilterBySmartModelId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance1 := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance1)
	assert.NoError(t, err)

	sminstance2 := &models.SmartModel{
		Name:       "Smart Phone",
		Identifier: "sp-001",
		Type:       "Device",
		Category:   "Mobile",
	}
	err = mokcedStorage.CreateSmartModel(sminstance2)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance1.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance2.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{SmartModelId: sminstance1.ID, Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}

func TestStorageListSmartFeatures_FilterByLimitAndOffset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	time.Sleep(1 * time.Microsecond)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{Limit: 1, Offset: 1})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}

func TestStorageListSmartFeatures_OrderByAscName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{OrderBy: "name asc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}

func TestStorageListSmartFeatures_OrderByDescName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{OrderBy: "name desc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(features))
	assert.Equal(t, feature2.Name, features[0].Name)
}

func TestStorageListSmartFeatures_OrderByAscCreatedAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{OrderBy: "created_at asc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}

func TestStorageListSmartFeatures_OrderByDescCreatedAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := mocks.MockConnection()
	defer cleanupMock()

	// Create models
	mokcedStorage, err := NewStorage(mockConnection)
	assert.NoError(t, err)

	// Create models
	sminstance := &models.SmartModel{
		Name:       "Smart Watch",
		Identifier: "sw-001",
		Type:       "Device",
		Category:   "Wearable",
	}
	err = mokcedStorage.CreateSmartModel(sminstance)
	assert.NoError(t, err)

	time.Sleep(1 * time.Microsecond)

	// Create features
	feature1 := &models.SmartFeature{
		Name:          "Future",
		Identifier:    "f-001",
		Functionality: "Predict the future",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature1)
	assert.NoError(t, err)

	feature2 := &models.SmartFeature{
		Name:          "Past",
		Identifier:    "f-002",
		Functionality: "Predict the past",
		SmartModelID:  sminstance.ID,
	}
	err = mokcedStorage.CreateSmartFeature(feature2)
	assert.NoError(t, err)

	// List features
	features, _, err := mokcedStorage.ListSmartFeatures(&filter.SmartFeatureFilter{OrderBy: "created_at desc", Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(features))
	assert.Equal(t, feature1.Name, features[0].Name)
}
