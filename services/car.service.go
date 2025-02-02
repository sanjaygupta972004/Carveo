package services

import (
	"carveo/models"
	"carveo/repositories"
	"carveo/utils"
	"errors"

	"github.com/gofrs/uuid"
)

type CarService interface {
	CreateCar(car models.Car, userID string) (models.Car, error)
	GetCarByID(id string) (models.Car, error)
	GetAllCars() ([]models.Car, error)
	GetCarByBrand(brand string, isEngine bool) (models.Car, error)
	UpdateCar(id string, car models.Car) (models.Car, error)
	DeleteCar(id string) error
}

type carService struct {
	carRepo repositories.CarRepository
}

func NewCarService(carRepo repositories.CarRepository) CarService {
	return &carService{
		carRepo: carRepo,
	}
}

func (s *carService) CreateCar(car models.Car, userID string) (models.Car, error) {
	if car.CarID != uuid.Nil {
		return models.Car{}, errors.New("engine_id should not be provided")
	}
	idUUID, err := utils.IsUUID(userID)
	if err != nil {
		return models.Car{}, err
	}
	data, err := s.carRepo.CreateCar(car, idUUID)

	if err != nil {
		return models.Car{}, err
	}

	return data, nil
}

func (s *carService) DeleteCar(id string) error {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return err
	}

	if err := s.carRepo.DeleteCar(idUUID); err != nil {
		return err
	}

	return nil
}

func (s *carService) GetAllCars() ([]models.Car, error) {
	var data []models.Car
	var err error
	data, err = s.carRepo.GetAllCars()
	if err != nil {
		return []models.Car{}, err
	}
	return data, nil
}

func (s *carService) GetCarByBrand(brand string, isEngine bool) (models.Car, error) {
	if brand == "" {
		err := errors.New("brand name is required to access car by brandName")
		return models.Car{}, err
	}
	data, err := s.carRepo.GetCarByBrand(brand, isEngine)
	if err != nil {
		return models.Car{}, err
	}
	return data, nil
}

func (s *carService) GetCarByID(id string) (models.Car, error) {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return models.Car{}, err
	}

	data, err := s.carRepo.GetCarByID(idUUID)
	if err != nil {
		return models.Car{}, err
	}

	return data, nil

}

func (s *carService) UpdateCar(id string, car models.Car) (models.Car, error) {
	idUUID, err := utils.IsUUID(id)
	if err != nil {
		return models.Car{}, err
	}
	data, err := s.carRepo.UpdateCar(idUUID, car)
	if err != nil {
		return models.Car{}, err
	}

	return data, nil
}
