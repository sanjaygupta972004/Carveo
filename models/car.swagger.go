package models

// CarSwagger represents the Swagger documentation for the Car model
// @Description Represents the details of a car
//
//	@Example {
//	  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
//	  "name": "Toyota Corolla",
//	  "year": "2021",
//	  "brand": "Toyota",
//	  "fuelType": "Petrol",
//	  "price": 20000.50,
//	  "createdAt": "2023-01-01T10:00:00Z",
//	  "updatedAt": "2023-02-01T12:00:00Z",
//	  "deletedAt": null
//	}
type CarSwagger struct {
	CarID     string  `json:"id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Name      string  `json:"name" example:"Tesla Model S"`
	Year      string  `json:"year" example:"2022"`
	Brand     string  `json:"brand" example:"Tesla"`
	FuelType  string  `json:"fuelType" example:"Electric"`
	Price     float64 `json:"price" example:"79999.99"`
	CreatedAt string  `json:"createdAt" example:"2025-01-01T12:00:00Z"`
	UpdatedAt string  `json:"updatedAt" example:"2025-01-01T12:00:00Z"`
}

// ErrorResponseCarSwagger defines the structure for error responses
type ErrorResponseCarSwagger struct {
	Message string      `json:"message" example:"Error message"`
	Status  string      `json:"status" example:"error"`
	Code    int         `json:"code" example:"400"`
	Error   interface{} `json:"error" example:"Error details"`
}

// SuccessResponseCarSwagger defines the structure for successful responses
type SuccessResponseCarSwagger struct {
	Message string      `json:"message" example:"Success message"`
	Status  string      `json:"status" example:"success"`
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data" example:"Response data details"`
}
