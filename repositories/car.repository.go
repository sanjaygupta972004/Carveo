package repositories

import "gorm.io/gorm"

type CarRepository interface {
	//
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{
		db: db,
	}
}
