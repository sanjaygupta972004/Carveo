package models

import (
	"carveo/validations"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Engine struct {
	EngineID      uuid.UUID `gorm:"type:uuid;primaryKey;unique; not null; index" json:"engineID"`
	Displacement  float64   `gorm:"type:decimal(10,2);not null" json:"displacement"`
	NoOfCylinders int       `gorm:"type:int;not null" json:"noOfCylinders"`
	CarRange      string    `gorm:"type:varchar(255);not null" json:"carRange"`

	// Relationship
	CarID uuid.UUID `gorm:"type:uuid;not null;index" json:"carID"`
}

func (e *Engine) BeforeCreate(tx *gorm.DB) error {
	e.EngineID = uuid.Must(uuid.NewV4())

	fmt.Println("Engine in model: ", e)

	err := validations.ValidateEngine(
		validations.Engine{
			Displacement:  e.Displacement,
			NoOfCylinders: e.NoOfCylinders,
			CarRange:      e.CarRange,
			EngineID:      e.EngineID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) TableName() string {
	return "Engines"
}
