package storage

import (
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
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
