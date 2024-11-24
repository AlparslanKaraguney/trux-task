package storage

import (
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/models"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

//go:generate mockgen -source=storage.go -destination=mocks/mock_storage.go -package=mocks
type Storage interface {
	CreateSmartModel(model *models.SmartModel) error
	GetSmartModel(id int32) (*models.SmartModel, error)
	UpdateSmartModel(model *models.SmartModel) error
	DeleteSmartModel(id int32) error

	CreateSmartFeature(feature *models.SmartFeature) error
	GetSmartFeature(id int32) (*models.SmartFeature, error)
	UpdateSmartFeature(feature *models.SmartFeature) error
	DeleteSmartFeature(id int32) error
}

type storageImpl struct {
	db *gorm.DB
}

func NewStorage(connection *gorm.DB) (Storage, error) {
	return &storageImpl{db: connection}, nil
}

func (s *storageImpl) CreateSmartModel(model *models.SmartModel) error {
	result := s.db.Create(model)
	return result.Error
}

func (s *storageImpl) GetSmartModel(id int32) (*models.SmartModel, error) {
	var model models.SmartModel
	result := s.db.Preload("Features").Where("id = ?", id).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &model, result.Error
}

func (s *storageImpl) UpdateSmartModel(model *models.SmartModel) error {
	if err := s.db.First(&models.SmartModel{}, model.ID).Error; err != nil {
		return ErrNotFound
	}

	result := s.db.Updates(model)

	return result.Error
}

func (s *storageImpl) DeleteSmartModel(id int32) error {
	if err := s.db.First(&models.SmartModel{}, id).Error; err != nil {
		return ErrNotFound
	}
	return s.db.Where("id = ?", id).Delete(&models.SmartModel{}).Error
}

func (s *storageImpl) CreateSmartFeature(feature *models.SmartFeature) error {
	result := s.db.Create(feature)
	return result.Error
}

func (s *storageImpl) GetSmartFeature(id int32) (*models.SmartFeature, error) {
	var feature models.SmartFeature
	result := s.db.Preload("SmartModel").
		Where("id = ?", id).First(&feature)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &feature, result.Error
}

func (s *storageImpl) UpdateSmartFeature(feature *models.SmartFeature) error {
	if err := s.db.First(&models.SmartFeature{}, feature.ID).Error; err != nil {
		return ErrNotFound
	}
	result := s.db.Save(feature)

	return result.Error
}

func (s *storageImpl) DeleteSmartFeature(id int32) error {
	if err := s.db.First(&models.SmartFeature{}, id).Error; err != nil {
		return ErrNotFound
	}
	return s.db.Where("id = ?", id).Delete(&models.SmartFeature{}).Error
}
