package repositories

import "gorm.io/gorm"

type EngineRepository interface {
}

type engineRepository struct {
	db *gorm.DB
}

func NewEngineRepository(db *gorm.DB) EngineRepository {
	return &engineRepository{
		db: db,
	}
}
