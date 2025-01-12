package services

import "carveo/repositories"

type EngineService interface {
}

type engineService struct {
	engineRepo repositories.EngineRepository
}

func NewEngineService(engineRepo repositories.EngineRepository) EngineService {
	return &engineService{
		engineRepo: engineRepo,
	}
}
