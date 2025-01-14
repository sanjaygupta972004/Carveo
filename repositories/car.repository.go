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
	err := r.db.Preload("engine").First(&car, id).Error
	return car, err
}

func (r *carRepository) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error
	return cars, err
}

func (r *carRepository) GetCarByBrand(brand string, isEngine bool) (models.Car, error) {
	var car models.Car
	if isEngine {
		err := r.db.Preload("engine").Where("brand = ?", brand).First(&car).Error
		return car, err
	}
	err := r.db.Where("brand = ?", brand).First(&car).Error
	return car, err
}

func (r *carRepository) CreateCar(car models.Car) (models.Car, error) {
	if err := car.BeforeCreate(r.db); err != nil {
		return models.Car{}, err
	}

	err := r.db.Create(&car).Error
	return car, err
}

func (r *carRepository) UpdateCar(id uuid.UUID, car models.Car) (models.Car, error) {
	err := r.db.Model(&models.Car{}).Where("car_id = ?", id).Updates(&car).Error
	return car, err
}

func (r *carRepository) DeleteCar(id uuid.UUID) error {
	err := r.db.Delete(&models.Car{}, id).Error
	return err
}
