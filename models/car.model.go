package models

import (
	"carveo/validations"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Car struct {
	CarID     uuid.UUID      `gorm:"type:uuid;primaryKey;not null; index; unique" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Year      string         `gorm:"type:varchar(4);not null" json:"year"`
	Brand     string         `gorm:"type:varchar(255);not null" json:"brand"`
	FuelType  string         `gorm:"type:varchar(255);not null" json:"fuelType"`
	Price     float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt time.Time      `json:"updatedAt"`
	UpdateAt  time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	// Relationship
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {
	c.CarID = uuid.Must(uuid.NewV4())
	return validations.ValidateCar(
		validations.Car{
			CarID:    c.CarID,
			Name:     c.Name,
			Year:     c.Year,
			Brand:    c.Brand,
			FuelType: c.FuelType,
		},
	)
}
func (c *Car) TableName() string {
	return "Cars"
}
