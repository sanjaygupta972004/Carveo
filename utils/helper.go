package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Error comparing password:", err)
		return fmt.Errorf("invalid password")
	}
	return nil
}
