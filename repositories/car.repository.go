package repositories

import (
	"carveo/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetCarByID(id uuid.UUID) (models.Car, error)
	GetAllCars() ([]models.Car, error)
	GetCarByBrand(brand string, isEngine bool) (models.Car, error)
	CreateCar(car models.Car) (models.Car, error)
	UpdateCar(id uuid.UUID, car models.Car) (models.Car, error)
	DeleteCar(id uuid.UUID) error
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{
		db: db,
	}
}

func (r *carRepository) GetCarByID(id uuid.UUID) (models.Car, error) {
	var car models.Car
	if err := r.db.Preload("engine").First(&car, id).Error; err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (r *carRepository) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	if err := r.db.Find(&cars).Error; err != nil {
		return []models.Car{}, err
	}
	return cars, nil
}

func (r *carRepository) GetCarByBrand(brand string, isEngine bool) (models.Car, error) {
	var car models.Car
	if isEngine {
		if err := r.db.Preload("engine").Where("brand = ?", brand).First(&car).Error; err != nil {
			return models.Car{}, err
		}
		return car, nil
	}
	if err := r.db.Where("brand = ?", brand).First(&car).Error; err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (r *carRepository) CreateCar(car models.Car) (models.Car, error) {
	if err := car.BeforeCreate(r.db); err != nil {
		return models.Car{}, err
	}

	if err := r.db.Where("car_id = ?", car.CarID).Find(&car).Error; err != nil {
		return models.Car{}, err
	}
	if err := r.db.Create(&car).Error; err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (r *carRepository) UpdateCar(id uuid.UUID, car models.Car) (models.Car, error) {
	var updatedCar models.Car
	if err := r.db.Where("car_id = ?", id).First(&updatedCar).Error; err != nil {
		return models.Car{}, err
	}

	if err := r.db.Model(&updatedCar).Where("car_id = ?", id).Updates(car).Error; err != nil {
		return models.Car{}, err
	}

	return updatedCar, nil
}

func (r *carRepository) DeleteCar(id uuid.UUID) error {
	if err := r.db.Delete(&models.Car{}, id).Error; err != nil {
		return err
	}
	return nil
}
