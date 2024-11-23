package storage

import (
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStorageCreateSmartModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection, cleanupMock := MockConnection()
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
