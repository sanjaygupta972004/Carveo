package controllers

import "carveo/services"

type EngineController interface {
}

type engineController struct {
	engineService services.EngineService
}

func NewEngineController(engineService services.EngineService) EngineController {
	return &engineController{
		engineService: engineService,
	}
}
