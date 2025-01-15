package repositories

import (
	"carveo/models"
	"fmt"

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
	var newEngine models.Engine
	if err := engine.BeforeCreate(r.db); err != nil {
		return models.Engine{}, err
	}

	var car models.Car
	if err := r.db.Where("car_id = ?", newEngine.CarID).First(&car).Error; err != nil {
		return models.Engine{}, err
	}

	if newEngine.CarID != car.CarID {
		return models.Engine{}, fmt.Errorf("car with ID %s not found", newEngine.CarID)
	}

	if err := r.db.Create(&newEngine).Error; err != nil {
		return models.Engine{}, err
	}
	return newEngine, nil
}

func (r *engineRepository) UpdateEngine(id uuid.UUID, engine models.Engine) (models.Engine, error) {
	var updatedEngine models.Engine
	if err := r.db.Where("engine_id = ?", id).First(&updatedEngine).Error; err != nil {
		return models.Engine{}, err
	}
	if err := r.db.Model(&updatedEngine).Where("engine_id = ?", id).Updates(engine).Error; err != nil {
		return models.Engine{}, err
	}
	return updatedEngine, nil
}

func (r *engineRepository) DeleteEngine(id uuid.UUID) error {
	if err := r.db.Delete(&models.Engine{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *engineRepository) GetEngineByID(id uuid.UUID) (models.Engine, error) {
	var engine models.Engine
	if err := r.db.First(&engine, id).Error; err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (r *engineRepository) GetAllEngines() ([]models.Engine, error) {
	var engines []models.Engine
	if err := r.db.Find(&engines).Error; err != nil {
		return []models.Engine{}, err
	}
	return engines, nil
}
