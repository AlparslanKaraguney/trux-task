package storage

import (
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/models"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

type Storage interface {
	CreateSmartModel(model *models.SmartModel) error
	GetSmartModel(identifier string) (*models.SmartModel, error)
	CreateSmartFeature(feature *models.SmartFeature) error
	GetSmartFeature(identifier string) (*models.SmartFeature, error)
}

type storageImpl struct {
	db *gorm.DB
}

func NewStorage(connection *gorm.DB) (Storage, error) {
	// Create a GORM Logrus Logger

	return &storageImpl{db: connection}, nil
}

func (s *storageImpl) CreateSmartModel(model *models.SmartModel) error {
	return s.db.Create(model).Error
}

func (s *storageImpl) GetSmartModel(identifier string) (*models.SmartModel, error) {
	var model models.SmartModel
	result := s.db.Preload("Features").Where("identifier = ?", identifier).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &model, result.Error
}

func (s *storageImpl) CreateSmartFeature(feature *models.SmartFeature) error {
	return s.db.Create(feature).Error
}

func (s *storageImpl) GetSmartFeature(identifier string) (*models.SmartFeature, error) {
	var feature models.SmartFeature
	result := s.db.Where("identifier = ?", identifier).First(&feature)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &feature, result.Error
}
