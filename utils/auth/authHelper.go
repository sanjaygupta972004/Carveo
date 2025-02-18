package auth

import (
	"carveo/config"
	"carveo/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SignJWTForUser(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.UserID,
		"email": user.Email,
		"iss":   "oauth-app-golang",
		"exp":   time.Now().Add(time.Hour * 24 * 2).Unix(),
	}
	secretKey := config.GetConfig().JwtSecret
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SignJWTRefreashTokenForUser(user *models.User) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.UserID,
		"email": user.Email,
		"iss":   "oauth-app-golang",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	secretKey := config.GetConfig().JwtSecret

	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedAccessToken, nil
}

func GenerateEmailVerificationToken(email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 15).Unix(), // 15-hours expiration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyEmailVerificationToken(tokenString string, secret string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(tokenString *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Error parsing token: ", err)
		return "", err
	}
	Email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("Error extracting email from token: ", err)
		return "", err
	}
	return Email, nil
}
