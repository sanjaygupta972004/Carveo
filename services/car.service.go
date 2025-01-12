package services

import "carveo/repositories"

type CarService interface {
}

type carService struct {
	carRepo repositories.CarRepository
}

func NewCarService(carRepo repositories.CarRepository) CarService {
	return &carService{
		carRepo: carRepo,
	}
}
