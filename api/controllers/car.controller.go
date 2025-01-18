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
	UpdateCar(*gin.Context)
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

type CarParameter struct {
	BrandName string `uri:"brandName" binding:"required"`
	IsEngine  *bool  `uri:"isEngine" binding:"required"`
}

func (cc *carController) CreateCar(ctx *gin.Context) {
	var carRequest models.Car
	if err := ctx.ShouldBindJSON(&carRequest); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Input fields should be json", err)
		return
	}

	data, err := cc.carService.CreateCar(carRequest)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}
	utils.SuccessResponse(ctx, http.StatusCreated, "Car successfully created", data)
}

func (cc *carController) GetAllCars(ctx *gin.Context) {
	data, err := cc.carService.GetAllCars()
	if err != nil {
		utils.SuccessResponse(ctx, http.StatusNotFound, utils.ErrNotFound, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All cars retrieved successfully", data)
}

func (cc *carController) GetCarByBrand(ctx *gin.Context) {
	var carParameter CarParameter
	if err := ctx.ShouldBindUri(&carParameter); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "All uri parameter required", err)
		return
	}
	if carParameter.IsEngine == nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "isEngine flag should be required", nil)
		return
	}
	data, err := cc.carService.GetCarByBrand(carParameter.BrandName, *carParameter.IsEngine)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, utils.ErrNotFound, err)
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "Retrieved data by brand name", data)
}

func (cc *carController) GetCarByID(ctx *gin.Context) {
	var uri struct {
		ID string `uri:"carID" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID should be required", err)
		return
	}

	data, err := cc.carService.GetCarByID(uri.ID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, utils.ErrNotFound, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Retrieved car by carID", data)

}

func (cc *carController) UpdateCar(ctx *gin.Context) {
	var carRequest models.Car
	if err := ctx.ShouldBindJSON(&carRequest); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Input fields should be json", err)
		return
	}
	var uri struct {
		ID string `uri:"carID" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID should be required", err)
		return
	}
	data, err := cc.carService.UpdateCar(uri.ID, carRequest)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Car successfully updated", data)
}

func (cc *carController) DeleteCar(ctx *gin.Context) {
	var uri struct {
		ID string `uri:"carID" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID should be required", err)
		return
	}
	err := cc.carService.DeleteCar(uri.ID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "Car successfully deleted", nil)
}
