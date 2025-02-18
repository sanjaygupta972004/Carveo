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

// @Summary Create Engine for a car
// @Description Create Engine for a Car by Car ID
// @Tags Engine
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param carID path string true "Car ID"
// @Param engine body models.EngineSwagger true "Engine Request"
// @Success 201 {object} models.SuccessResponseEngineSwagger "Engine created successfully"
// @Failure 400 {object} models.ErrorResponseEngineSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseEngineSwagger "Internal server error"
// @Router /engiens/createEngine{carID}/ [post]
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

// @Summary Get Engine by ID
// @Description Get Engine by Engine ID
// @Tags Engine
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param engineID path string true "Engine ID"
// @Success 200 {object} models.SuccessResponseEngineSwagger "Engine retrieved successfully"
// @Failure 500 {object} models.ErrorResponseEngineSwagger "Internal server error"
// @Router /engines/getEngineByID/{engineID} [get]
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

// @Summary Get All Engines
// @Description Endpoint to get all engines
// @Tags Engine
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} []models.SuccessResponseEngineSwagger "All engines retrieved successfully"
// @Failure 500 {object} models.ErrorResponseEngineSwagger "Internal server error"
// @Router /engines/getAllEngines [get]
func (ec *engineController) GetEngines(c *gin.Context) {
	engineData, err := ec.engineService.GetAllEngines()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Engines retrieved successfully", engineData)
}

// @Summary Update Engine
// @Description Update Engine by Engine ID
// @Tags Engine
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param engineID path string true "Engine ID"
// @Param engine body models.EngineSwagger true "Engine Request"
// @Success 200 {object} models.SuccessResponseEngineSwagger "Engine updated successfully"
// @Failure 400 {object} models.ErrorResponseEngineSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseEngineSwagger "Internal server error"
// @Router /engines/updateEngine/{engineID} [patch]
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

// @Summary Delete Engine
// @Description Delete Engine by Engine ID
// @Tags Engine
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param engineID path string true "Engine ID"
// @Success 200 {string} string "Engine deleted successfully"
// @Failure 500 {object} models.ErrorResponseEngineSwagger "Internal server error"
// @Router /engines/deleteEngine/{engineID} [delete]
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
