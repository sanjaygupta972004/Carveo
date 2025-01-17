package validations

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type Engine struct {
	EngineID      uuid.UUID
	Displacement  float64
	NoOfCylinders int
	CarRange      string
}

func ValidateEngine(eng Engine) error {
	fmt.Println("Engine in validator: ", eng)
	if eng.Displacement <= 0 {
		return fmt.Errorf("displacement must be greater than 0")
	}
	if eng.EngineID == uuid.Nil {
		return fmt.Errorf("engine ID is required")
	}
	if eng.NoOfCylinders <= 0 {
		return fmt.Errorf("no of cylinders must be greater than 0")
	}
	if eng.CarRange == "" {
		return fmt.Errorf("car range is required")
	}
	return nil
}
