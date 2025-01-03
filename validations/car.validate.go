package validations

import (
	"errors"

	"github.com/gofrs/uuid"
)

type Car struct {
	CarID    uuid.UUID
	Name     string
	Year     string
	Brand    string
	FuelType string
	Price    float64
}

func contains(fuels []string, item string) bool {
	for _, s := range fuels {
		if s == item {
			return true
		}
	}
	return false
}

func ValidateCar(car Car) error {
	fuelType := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	if car.Name == "" {
		return errors.New("Name must not be empty")
	}
	if car.Year == "" {
		return errors.New("Year must not be empty")
	}
	if car.Brand == "" {
		return errors.New("Brand must not be empty")
	}
	if car.CarID == uuid.Nil {
		return errors.New("CarID must not be empty")
	}
	if car.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}
	if !contains(fuelType, car.FuelType) {
		return errors.New("Fuel type must be Petrol, Diesel, Electric, or Hybrid")
	}

	return nil
}
