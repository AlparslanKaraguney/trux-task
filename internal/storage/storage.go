package storage

import (
	"errors"

	"github.com/AlparslanKaraguney/trux-task/internal/entities"
	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/internal/storage/filter"
	apperrors "github.com/AlparslanKaraguney/trux-task/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"

	"gorm.io/gorm"
)

//go:generate mockgen -source=storage.go -destination=mocks/mock_storage.go -package=mocks
type Storage interface {
	CreateSmartModel(model *models.SmartModel) error
	GetSmartModel(id int32) (*models.SmartModel, error)
	UpdateSmartModel(model *models.SmartModel) error
	DeleteSmartModel(id int32) error
	ListSmartModels(filter *filter.SmartModelFilter) ([]models.SmartModel, *entities.Pagination, error)

	CreateSmartFeature(feature *models.SmartFeature) error
	GetSmartFeature(id int32) (*models.SmartFeature, error)
	UpdateSmartFeature(feature *models.SmartFeature) error
	DeleteSmartFeature(id int32) error
	ListSmartFeatures(filter *filter.SmartFeatureFilter) ([]models.SmartFeature, *entities.Pagination, error)
}

type storageImpl struct {
	db *gorm.DB
}

func NewStorage(connection *gorm.DB) (Storage, error) {
	return &storageImpl{db: connection}, nil
}

func (s *storageImpl) CreateSmartModel(model *models.SmartModel) error {
	result := s.db.Create(model)
	var pgErr *pgconn.PgError
	if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
		return apperrors.ErrAlreadyExists
	}
	return result.Error
}

func (s *storageImpl) GetSmartModel(id int32) (*models.SmartModel, error) {
	var model models.SmartModel
	result := s.db.Preload("Features").Where("id = ?", id).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrNotFound
	}
	return &model, result.Error
}

func (s *storageImpl) UpdateSmartModel(model *models.SmartModel) error {
	if err := s.db.First(&models.SmartModel{}, model.ID).Error; err != nil {
		return apperrors.ErrNotFound
	}

	result := s.db.Updates(model)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
			return apperrors.ErrAlreadyExists
		}
	}

	return result.Error
}

func (s *storageImpl) DeleteSmartModel(id int32) error {
	if err := s.db.First(&models.SmartModel{}, id).Error; err != nil {
		return apperrors.ErrNotFound
	}
	return s.db.Where("id = ?", id).Delete(&models.SmartModel{}).Error
}

func (s *storageImpl) ListSmartModels(filter *filter.SmartModelFilter) ([]models.SmartModel, *entities.Pagination, error) {
	var models []models.SmartModel
	query := s.db.Model(&models)
	if filter != nil {
		if filter.Identifier != "" {
			query = query.Where("identifier = ?", filter.Identifier)
		}
		if filter.Name != "" {
			query = query.Where("name = ?", filter.Name)
		}
		if filter.Type != "" {
			query = query.Where("type = ?", filter.Type)
		}
		if filter.Category != "" {
			query = query.Where("category = ?", filter.Category)
		}
	}

	switch filter.OrderBy {
	case "name asc":
		query = query.Order("name asc")
	case "name desc":
		query = query.Order("name desc")
	case "created_at asc":
		query = query.Order("created_at asc")
	case "created_at desc":
		query = query.Order("created_at desc")
	case "":
		query = query.Order("created_at desc")
	default:
		return nil, nil, apperrors.ErrInvalidOrderBy
	}

	query, pagination := paginate(int(filter.Limit), int(filter.Offset), query)
	result := query.Preload("Features").Find(&models)
	return models, pagination, result.Error
}

func (s *storageImpl) CreateSmartFeature(feature *models.SmartFeature) error {
	result := s.db.Create(feature)
	var pgErr *pgconn.PgError
	if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
		return apperrors.ErrAlreadyExists
	}

	return result.Error
}

func (s *storageImpl) GetSmartFeature(id int32) (*models.SmartFeature, error) {
	var feature models.SmartFeature
	result := s.db.Preload("SmartModel").
		Where("id = ?", id).First(&feature)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrNotFound
	}
	return &feature, result.Error
}

func (s *storageImpl) UpdateSmartFeature(feature *models.SmartFeature) error {
	if err := s.db.First(&models.SmartFeature{}, feature.ID).Error; err != nil {
		return apperrors.ErrNotFound
	}
	result := s.db.Save(feature)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
			return apperrors.ErrAlreadyExists
		}
	}

	return result.Error
}

func (s *storageImpl) DeleteSmartFeature(id int32) error {
	if err := s.db.First(&models.SmartFeature{}, id).Error; err != nil {
		return apperrors.ErrNotFound
	}
	return s.db.Where("id = ?", id).Delete(&models.SmartFeature{}).Error
}

func (s *storageImpl) ListSmartFeatures(filter *filter.SmartFeatureFilter) ([]models.SmartFeature, *entities.Pagination, error) {
	var features []models.SmartFeature
	query := s.db.Model(&features)
	if filter != nil {
		if filter.Identifier != "" {
			query = query.Where("identifier = ?", filter.Identifier)
		}
		if filter.Name != "" {
			query = query.Where("name = ?", filter.Name)
		}
		if filter.Functionality != "" {
			query = query.Where("functionality = ?", filter.Functionality)
		}
		if filter.SmartModelId != 0 {
			query = query.Where("smart_model_id = ?", filter.SmartModelId)
		}
	}

	switch filter.OrderBy {
	case "name asc":
		query = query.Order("name asc")
	case "name desc":
		query = query.Order("name desc")
	case "created_at asc":
		query = query.Order("created_at asc")
	case "created_at desc":
		query = query.Order("created_at desc")
	case "":
		query = query.Order("created_at desc")
	default:
		return nil, nil, apperrors.ErrInvalidOrderBy
	}

	query, pagination := paginate(int(filter.Limit), int(filter.Offset), query)
	result := query.Find(&features)
	return features, pagination, result.Error
}
