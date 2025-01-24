package controllers

import (
	"carveo/models"
	"carveo/services"
	"carveo/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type EngineController interface {
	CreateEngine(*gin.Context)
	GetEngine(*gin.Context)
	GetEngines(*gin.Context)
	UpdateEngine(*gin.Context)
	DeleteEngine(*gin.Context)
}

func NewEngineController(engineService services.EngineService) EngineController {
	return &engineController{
		engineService: engineService,
	}
}

type engineController struct {
	engineService services.EngineService
}

func (ec *engineController) CreateEngine(c *gin.Context) {
	carID := c.Param("carID")
	if carID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Car ID is required", nil)
		return
	}
	var engineRequest models.Engine
	if err := c.ShouldBindJSON(&engineRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	engine, err := ec.engineService.CreateEngine(engineRequest, carID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Engine created successfully", engine)
}

func (ec *engineController) GetEngine(c *gin.Context) {
	engineID := c.Param("engineID")
	if engineID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Engine ID is required", nil)
		return
	}
	engine, err := ec.engineService.GetEngineById(engineID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Engine retrieved successfully", engine)
}

func (ec *engineController) GetEngines(c *gin.Context) {
	engineData, err := ec.engineService.GetAllEngines()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Engines retrieved successfully", engineData)
}

func (ec *engineController) UpdateEngine(c *gin.Context) {
	var engineRequest models.Engine
	if err := c.ShouldBindJSON(&engineRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	engineID := c.Param("engineID")
	if engineID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Engine ID is required", nil)
		return
	}
	engine, err := ec.engineService.UpdateEngine(engineID, engineRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Engine updated successfully", engine)
}

func (ec *engineController) DeleteEngine(c *gin.Context) {
	engineID := c.Param("engineID")
	if engineID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Engine ID is required", nil)
		return
	}
	err := ec.engineService.DeleteEngine(engineID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Engine deleted successfully", nil)
}
