package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func IsUUID(id string) (uuid.UUID, error) {
	if id == "" {
		err := errors.New("ID is required to check is uuid type or not")
		return uuid.UUID{}, err
	}
	idUUID, err := uuid.FromString(id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return idUUID, nil
}

func ErrorResponse(ctx *gin.Context, statusCode int, customMessage string, details interface{}) {
	if details != nil {
		log.Printf("Error : %v", details)
	}
	fmt.Println("Error in helper function : ", details)
	errorDetails := fmt.Sprintf("%v", details)
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"message": customMessage,
		"error":   errorDetails,
	})
}

func SuccessResponse(ctx *gin.Context, statusCode int, customMessage string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": customMessage,
		"data":    data,
	})

}
