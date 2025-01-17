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
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	// Relationship
	Engine Engine `gorm:"foreignKey:CarID;references:CarID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (c *Car) BeforeCreate(tx *gorm.DB) error {
	c.CarID = uuid.Must(uuid.NewV4())
	err := validations.ValidateCar(
		validations.Car{
			CarID:    c.CarID,
			Name:     c.Name,
			Year:     c.Year,
			Brand:    c.Brand,
			FuelType: c.FuelType,
			Price:    c.Price,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func (c *Car) TableName() string {
	return "Cars"
}
