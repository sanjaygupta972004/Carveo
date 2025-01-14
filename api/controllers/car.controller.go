package controllers

import (
	"carveo/models"
	"carveo/services"
	"carveo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarController interface {
	CreateCar(*gin.Context)
	GetAllCars(*gin.Context)
	GetCarByID(*gin.Context)
	GetCarByBrand(*gin.Context)
	UpdataCar(*gin.Context)
	DeleteCar(*gin.Context)
}

type carController struct {
	carService services.CarService
}

func NewCarController(carService services.CarService) CarController {
	return &carController{
		carService: carService,
	}
}

func (cc *carController) CreateCar(ctx *gin.Context) {
	var carRequest models.Car
	if err := ctx.ShouldBindJSON(carRequest); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Input fields should be json", err)
	}

	data, err := cc.carService.CreateCar(carRequest)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Car successfully created", data)

}

func (cc *carController) GetAllCars(ctx *gin.Context) {

}

func (cc *carController) GetCarByBrand(ctx *gin.Context) {

}

func (cc *carController) GetCarByID(ctx *gin.Context) {

}

func (cc *carController) UpdataCar(ctx *gin.Context) {

}

func (cc *carController) DeleteCar(ctx *gin.Context) {

}
