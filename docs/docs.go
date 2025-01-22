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
            "email": "support@carveo.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Check if the server is running and healthy.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
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
	Host:             "localhost:8090",
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
