package models

// CarSwagger represents the Swagger documentation for the Car model
// @Description Represents the details of a car
type CarSwagger struct {
	CarID     string  `json:"id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Name      string  `json:"name" example:"Tesla Model S"`
	Year      string  `json:"year" example:"2022"`
	Brand     string  `json:"brand" example:"Tesla"`
	FuelType  string  `json:"fuelType" example:"Electric"`
	Price     float64 `json:"price" example:"79999.99"`
	UserID    string  `json:"userID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	CreatedAt string  `json:"createdAt" example:"2025-01-01T12:00:00Z"`
	UpdatedAt string  `json:"updatedAt" example:"2025-01-01T12:00:00Z"`
}

// ErrorResponseCarSwagger defines the structure for error responses
type ErrorResponseCarSwagger struct {
	Message string `json:"message" example:"Error message"`
	Status  string `json:"status" example:"error"`
	Code    int    `json:"code" example:"400"`
	Error   string `json:"error" example:"Detailed error message"`
}

// SuccessResponseCarSwagger defines the structure for successful responses
type SuccessResponseCarSwagger struct {
	Message string `json:"message" example:"Success message"`
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"200"`
	Data    struct {
		CarID     string  `json:"id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
		Name      string  `json:"name" example:"Tesla Model S"`
		Year      string  `json:"year" example:"2022"`
		Brand     string  `json:"brand" example:"Tesla"`
		FuelType  string  `json:"fuelType" example:"Electric"`
		Price     float64 `json:"price" example:"79999.99"`
		UserID    string  `json:"userID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
		CreatedAt string  `json:"createdAt" example:"2025-01-01T12:00:00Z"`
		UpdatedAt string  `json:"updatedAt" example:"2025-01-01T12:00:00Z"`
	} `json:"data"`
}

type CarResponseByGetByBrandIsEngine struct {
	Car    CarSwagger
	Engine struct {
		EngineID      string  `json:"engineID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
		Displacement  float64 `json:"displacement" example:"3.5"`
		NoOfCylinders int     `json:"noOfCylinders" example:"6"`
		CarRange      string  `json:"carRange" example:"300-400 miles"`
		CarID         string  `json:"carID" example:"a123c456-78de-90fg-1234-h567ijkl8901"`
	} `json:"engine"`
}

type CarResponseByGetByID struct {
	Car    SuccessResponseCarSwagger
	Engine struct {
		EngineID      string  `json:"engineID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
		Displacement  float64 `json:"displacement" example:"3.5"`
		NoOfCylinders int     `json:"noOfCylinders" example:"6"`
		CarRange      string  `json:"carRange" example:"300-400 miles"`
		CarID         string  `json:"carID" example:"a123c456-78de-90fg-1234-h567ijkl8901"`
	} `json:"engine"`
}
