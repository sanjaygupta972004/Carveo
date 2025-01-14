package repositories

import (
	"carveo/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type EngineRepository interface {
	CreateEngine(engine models.Engine) (models.Engine, error)
	UpdateEngine(id uuid.UUID, engine models.Engine) (models.Engine, error)
	DeleteEngine(id uuid.UUID) error
	GetEngineByID(id uuid.UUID) (models.Engine, error)
	GetAllEngines() ([]models.Engine, error)
}

type engineRepository struct {
	db *gorm.DB
}

func NewEngineRepository(db *gorm.DB) EngineRepository {
	return &engineRepository{
		db: db,
	}
}

func (r *engineRepository) CreateEngine(engine models.Engine) (models.Engine, error) {
	return engine, nil
}

func (r *engineRepository) UpdateEngine(id uuid.UUID, engine models.Engine) (models.Engine, error) {
	return engine, nil
}

func (r *engineRepository) DeleteEngine(id uuid.UUID) error {
	return nil
}

func (r *engineRepository) GetEngineByID(id uuid.UUID) (models.Engine, error) {
	return models.Engine{}, nil
}

func (r *engineRepository) GetAllEngines() ([]models.Engine, error) {
	return []models.Engine{}, nil
}
