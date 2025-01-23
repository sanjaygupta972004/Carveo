package models

// EngineSwagger represents the Swagger documentation for the Engine model
// @Description Represents the engine details of a car
//
//	@Example {
//	  "engineID": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
//	  "displacement": 3.5,
//	  "noOfCylinders": 6,
//	  "carRange": "300-400 miles",
//	  "carID": "a123c456-78de-90fg-1234-h567ijkl8901"
//	}
type EngineSwagger struct {
	EngineID      string  `json:"engineID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Displacement  float64 `json:"displacement" example:"3.5"`
	NoOfCylinders int     `json:"noOfCylinders" example:"6"`
	CarRange      string  `json:"carRange" example:"300-400 miles"`
	CarID         string  `json:"carID" example:"a123c456-78de-90fg-1234-h567ijkl8901"`
}

// ErrorResponseEngineSwagger defines the structure for error responses
type ErrorResponseEngineSwagger struct {
	Message string      `json:"message" example:"Error message "`
	Status  string      `json:"status" example:"error or failed to validation request data or server error"`
	Code    int         `json:"code" example:"400 or 500 or any other status code"`
	Error   interface{} `json:"error" example:"Error details"`
}

// SuccessResponseEngineSwagger defines the structure for successful responses
type SuccessResponseEngineSwagger struct {
	Message string      `json:"message" example:"Success message "`
	Status  string      `json:"status" example:"success"`
	Code    int         `json:"code" example:"200 or 201 or any other status code"`
	Data    interface{} `json:"data" example:" Resonse Data details"`
}
