package validations

import (
	"errors"

	"github.com/gofrs/uuid"
)

type Engine struct {
	EngineID      uuid.UUID
	Displacement  float64
	NoOfCylinders int
	CarRange      string
}

func ValidateEngine(eng Engine) error {
	if eng.Displacement <= 0 {
		return errors.New("Displacement must be greater than 0")
	}
	if eng.EngineID == uuid.Nil {
		return errors.New("EngineID must not be empty")
	}
	if eng.NoOfCylinders <= 0 {
		return errors.New("Number of cylinders must be greater than 0")
	}
	if eng.CarRange == "" {
		return errors.New("Car range must not be empty")
	}
	return nil
}
