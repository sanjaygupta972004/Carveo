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

// @Summary Create a new car
// @Description Create a new car with the input payload. This endpoint allows you to add a car to the database.
// @Tags Car
// @Accept json
// @Produce json
// @Param car body models.CarSwagger true "Car Request"
// @Success 201 {object} models.SuccessResponseCarSwagger "Car created successfully"
// @Failure 400 {object} models.ErrorResponseCarSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseCarSwagger  "Internal server error"
// @Router /cars/createCar [post]
func (cc *carController) CreateCar(ctx *gin.Context) {
	var carRequest models.Car
	if err := ctx.ShouldBindJSON(&carRequest); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Input fields should be json", err)
		return
	}

	useID, err := utils.GetUserIdFromHeader(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "User ID should be provided", err)
		return
	}
	data, err := cc.carService.CreateCar(carRequest, useID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}
	utils.SuccessResponse(ctx, http.StatusCreated, "Car successfully created", data)
}

// @Summary Get all cars
// @Description Retrieve all cars from the database
// @Tags Car
// @Produce json
// @Success 200 {object} []models.SuccessResponseCarSwagger "All cars retrieved successfully"
// @Failure 404 {object} models.ErrorResponseCarSwagger "No cars found"
// @Failure 500 {object} models.ErrorResponseCarSwagger  "Internal server error"
// @Router /cars/getAllCars [get]
func (cc *carController) GetAllCars(ctx *gin.Context) {
	data, err := cc.carService.GetAllCars()
	if err != nil {
		utils.SuccessResponse(ctx, http.StatusNotFound, utils.ErrNotFound, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All cars retrieved successfully", data)
}

// @Summary Get car by brand name and engine type
// @Description Retrieve car by brand name and engine type from the database
// @Tags Car
// @Produce json
// @Param brandName path string true "Brand name of the car to retrieve"
// @Param isEngine path boolean true "Engine type of the car"
// @Success 200 {object} models.CarResponseByGetByBrandIsEngine "Car retrieved successfully by brand name and engine type"
// @success 200 {object} models.SuccessResponseCarSwagger "Car retrieved successfully by only brand name without engine type"
// @Failure 404 {object} models.ErrorResponseCarSwagger "No car found"
// @Failure 500 {object} models.ErrorResponseCarSwagger  "Internal server error"
// @Router /cars/getCarByBrand/{brandName}/{isEngine} [get]
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

// @Summary Get car by ID
// @Description Retrieve car by ID from the database
// @Tags Car
// @Produce json
// @Param carID path string true "ID of the car to retrieve"
// @Success 200 {object} models.CarResponseByGetByID "Car retrieved successfully by ID"
// @Failure 404 {object} models.ErrorResponseCarSwagger "No car found"
// @Failure 500 {object} models.ErrorResponseCarSwagger  "Internal server error"
// @Router /cars/getCarByID/{carID} [get]
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

// @Summary Update a car
// @Description Update a car with the input payload. This endpoint allows you to update a car in the database.
// @Tags Car
// @Accept json
// @Produce json
// @Param carID path string true "ID of the car to update"
// @Param car body models.CarSwagger true "Car Request"
// @Success 200 {object} models.SuccessResponseCarSwagger "Car updated successfully"
// @Failure 400 {object} models.ErrorResponseCarSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseCarSwagger  "Internal server error"
// @Router /cars/updateCar/{carID} [patch]
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

// @Summary Delete a car
// @Description Delete a car with the input payload. This endpoint allows you to delete a car in the database.
// @Tags Car
// @Accept json
// @Produce json
// @Param carID path string true "ID of the car to delete"
// @Success 200 {string} string "Car deleted successfully"
// @Failure 400 {object} models.ErrorResponseCarSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseCarSwagger "Internal server error"
// @Router /cars/deleteCar/{carID} [delete]
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
