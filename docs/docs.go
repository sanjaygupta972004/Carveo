// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://carveo.com/terms",
        "contact": {
            "name": "Support Team",
            "url": "https://carveo.com/support",
            "email": "sanjaygupta07054@gmailcom"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars/createCar": {
            "post": {
                "description": "Create a new car with the input payload. This endpoint allows you to add a car to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Create a new car",
                "parameters": [
                    {
                        "description": "Car Request",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CarSwagger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Car created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseCarSwagger"
                        }
                    },
                    "400": {
                        "description": "Invalid input fields or JSON format",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/cars/deleteCar/{carID}": {
            "delete": {
                "description": "Delete a car with the input payload. This endpoint allows you to delete a car in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Delete a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the car to delete",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input fields or JSON format",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/cars/getAllCars": {
            "get": {
                "description": "Retrieve all cars from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Get all cars",
                "responses": {
                    "200": {
                        "description": "All cars retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SuccessResponseCarSwagger"
                            }
                        }
                    },
                    "404": {
                        "description": "No cars found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/cars/getCarByBrand/{brandName}/{isEngine}": {
            "get": {
                "description": "Retrieve car by brand name and engine type from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Get car by brand name and engine type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Brand name of the car to retrieve",
                        "name": "brandName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Engine type of the car",
                        "name": "isEngine",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car retrieved successfully by only brand name without engine type",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseCarSwagger"
                        }
                    },
                    "404": {
                        "description": "No car found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/cars/getCarByID/{carID}": {
            "get": {
                "description": "Retrieve car by ID from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Get car by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the car to retrieve",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car retrieved successfully by ID",
                        "schema": {
                            "$ref": "#/definitions/models.CarResponseByGetByID"
                        }
                    },
                    "404": {
                        "description": "No car found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/cars/updateCar/{carID}": {
            "patch": {
                "description": "Update a car with the input payload. This endpoint allows you to update a car in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Update a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the car to update",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Car Request",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CarSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseCarSwagger"
                        }
                    },
                    "400": {
                        "description": "Invalid input fields or JSON format",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseCarSwagger"
                        }
                    }
                }
            }
        },
        "/engiens/createEngine{carID}/": {
            "post": {
                "description": "Create Engine for a Car by Car ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Engine"
                ],
                "summary": "Create Engine for a Car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Engine Request",
                        "name": "engine",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EngineSwagger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Engine created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseEngineSwagger"
                        }
                    },
                    "400": {
                        "description": "Invalid input fields or JSON format",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    }
                }
            }
        },
        "/engines/deleteEngine/{engineID}": {
            "delete": {
                "description": "Delete Engine by Engine ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Engine"
                ],
                "summary": "Delete Engine",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Engine ID",
                        "name": "engineID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Engine deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    }
                }
            }
        },
        "/engines/getAllEngines": {
            "get": {
                "description": "Endpoint to get all engines",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Engine"
                ],
                "summary": "Get All Engines",
                "responses": {
                    "200": {
                        "description": "All engines retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SuccessResponseEngineSwagger"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    }
                }
            }
        },
        "/engines/getEngineByID/{engineID}": {
            "get": {
                "description": "Get Engine by Engine ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Engine"
                ],
                "summary": "Get Engine by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Engine ID",
                        "name": "engineID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Engine retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseEngineSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    }
                }
            }
        },
        "/engines/updateEngine/{engineID}": {
            "patch": {
                "description": "Update Engine by Engine ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Engine"
                ],
                "summary": "Update Engine",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Engine ID",
                        "name": "engineID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Engine Request",
                        "name": "engine",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EngineSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Engine updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponseEngineSwagger"
                        }
                    },
                    "400": {
                        "description": "Invalid input fields or JSON format",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponseEngineSwagger"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Provides a simple health check to verify the server's operational status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check Endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routers.HealthCheckResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CarResponseByGetByBrandIsEngine": {
            "type": "object",
            "properties": {
                "car": {
                    "$ref": "#/definitions/models.CarSwagger"
                },
                "engine": {
                    "type": "object",
                    "properties": {
                        "carID": {
                            "type": "string",
                            "example": "a123c456-78de-90fg-1234-h567ijkl8901"
                        },
                        "carRange": {
                            "type": "string",
                            "example": "300-400 miles"
                        },
                        "displacement": {
                            "type": "number",
                            "example": 3.5
                        },
                        "engineID": {
                            "type": "string",
                            "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                        },
                        "noOfCylinders": {
                            "type": "integer",
                            "example": 6
                        }
                    }
                }
            }
        },
        "models.CarResponseByGetByID": {
            "type": "object",
            "properties": {
                "car": {
                    "$ref": "#/definitions/models.SuccessResponseCarSwagger"
                },
                "engine": {
                    "type": "object",
                    "properties": {
                        "carID": {
                            "type": "string",
                            "example": "a123c456-78de-90fg-1234-h567ijkl8901"
                        },
                        "carRange": {
                            "type": "string",
                            "example": "300-400 miles"
                        },
                        "displacement": {
                            "type": "number",
                            "example": 3.5
                        },
                        "engineID": {
                            "type": "string",
                            "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                        },
                        "noOfCylinders": {
                            "type": "integer",
                            "example": 6
                        }
                    }
                }
            }
        },
        "models.CarSwagger": {
            "description": "Represents the details of a car",
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string",
                    "example": "Tesla"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2025-01-01T12:00:00Z"
                },
                "fuelType": {
                    "type": "string",
                    "example": "Electric"
                },
                "id": {
                    "type": "string",
                    "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                },
                "name": {
                    "type": "string",
                    "example": "Tesla Model S"
                },
                "price": {
                    "type": "number",
                    "example": 79999.99
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2025-01-01T12:00:00Z"
                },
                "year": {
                    "type": "string",
                    "example": "2022"
                }
            }
        },
        "models.EngineSwagger": {
            "description": "Represents the details of a Enigne",
            "type": "object",
            "properties": {
                "carID": {
                    "type": "string",
                    "example": "a123c456-78de-90fg-1234-h567ijkl8901"
                },
                "carRange": {
                    "type": "string",
                    "example": "300-400 miles"
                },
                "displacement": {
                    "type": "number",
                    "example": 3.5
                },
                "engineID": {
                    "type": "string",
                    "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                },
                "noOfCylinders": {
                    "type": "integer",
                    "example": 6
                }
            }
        },
        "models.ErrorResponseCarSwagger": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "error": {
                    "type": "string",
                    "example": "Detailed error message"
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "models.ErrorResponseEngineSwagger": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "error": {
                    "type": "string",
                    "example": "Error details"
                },
                "message": {
                    "type": "string",
                    "example": "Error message "
                },
                "status": {
                    "type": "string",
                    "example": "error or failed to validation request data or server error"
                }
            }
        },
        "models.SuccessResponseCarSwagger": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "brand": {
                            "type": "string",
                            "example": "Tesla"
                        },
                        "createdAt": {
                            "type": "string",
                            "example": "2025-01-01T12:00:00Z"
                        },
                        "fuelType": {
                            "type": "string",
                            "example": "Electric"
                        },
                        "id": {
                            "type": "string",
                            "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                        },
                        "name": {
                            "type": "string",
                            "example": "Tesla Model S"
                        },
                        "price": {
                            "type": "number",
                            "example": 79999.99
                        },
                        "updatedAt": {
                            "type": "string",
                            "example": "2025-01-01T12:00:00Z"
                        },
                        "year": {
                            "type": "string",
                            "example": "2022"
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Success message"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "models.SuccessResponseEngineSwagger": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "carID": {
                            "type": "string",
                            "example": "a123c456-78de-90fg-1234-h567ijkl8901"
                        },
                        "carRange": {
                            "type": "string",
                            "example": "300-400 miles"
                        },
                        "displacement": {
                            "type": "number",
                            "example": 3.5
                        },
                        "engineID": {
                            "type": "string",
                            "example": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
                        },
                        "noOfCylinders": {
                            "type": "integer",
                            "example": 6
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Success message "
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "routers.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "error": {},
                "message": {
                    "type": "string",
                    "example": "Server is running and healthy"
                },
                "status": {
                    "type": "string",
                    "example": "healthy"
                }
            }
        }
    },
    "securityDefinitions": {
        "CarveoAPIKey": {
            "description": "JWT token required to access endpoints. Add the token in the \"Authorization\" header in the format: Bearer \u003cJWT_TOKEN\u003e.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "Endpoints related to car operations.",
            "name": "Car Management"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "http://localhost:8090",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Carevo Car Management API Documentation",
	Description:      "API documentation for Carevo, a car management system for tracking, servicing, and managing vehicles.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
