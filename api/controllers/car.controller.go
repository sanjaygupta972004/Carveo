package controllers

import "carveo/services"

type CarController interface {
	CreateCar()
	GetAllCars()
	GetCarByID()
	GetCarByBrand()
	UpdataCar()
	DeleteCar()
}

type carRequest struct {
}

type carController struct {
	carService services.CarService
}

func NewCarController(carService services.CarService) CarController {
	return &carController{
		carService: carService,
	}
}
