package services

import (
	"carveo/models"
	"carveo/repositories"
	"carveo/utils"
	"errors"

	"github.com/gofrs/uuid"
)

type EngineService interface {
	GetAllEngines() ([]models.Engine, error)
	GetEngineById(id string) (models.Engine, error)
	CreateEngine(engine models.Engine, carID string) (models.Engine, error)
	UpdateEngine(id string, engine models.Engine) (models.Engine, error)
	DeleteEngine(id string) error
}

type engineService struct {
	engineRepo repositories.EngineRepository
}

func NewEngineService(engineRepo repositories.EngineRepository) EngineService {
	return &engineService{
		engineRepo: engineRepo,
	}
}

func (s *engineService) GetAllEngines() ([]models.Engine, error) {
	return s.engineRepo.GetAllEngines()
}

func (s *engineService) GetEngineById(id string) (models.Engine, error) {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return models.Engine{}, err
	}
	return s.engineRepo.GetEngineByID(idUUID)
}

func (s *engineService) CreateEngine(engine models.Engine, carID string) (models.Engine, error) {
	idUUID, err := utils.IsUUID(carID)
	if err != nil {
		return models.Engine{}, err
	}
	if idUUID == uuid.Nil {
		return models.Engine{}, errors.New("Car ID is required")
	}

	if idUUID != uuid.Nil {
		engine.CarID = idUUID
	}
	return s.engineRepo.CreateEngine(engine)
}

func (s *engineService) UpdateEngine(id string, engine models.Engine) (models.Engine, error) {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return models.Engine{}, err
	}
	return s.engineRepo.UpdateEngine(idUUID, engine)
}

func (s *engineService) DeleteEngine(id string) error {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return err
	}
	return s.engineRepo.DeleteEngine(idUUID)
}
