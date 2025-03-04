package controllers

import (
	"carveo/config"
	"carveo/models"
	"carveo/services"
	"carveo/utils"
	"carveo/utils/auth"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GenerateResetPasswordToken(c *gin.Context)
	ValidateResetPasswordToken(c *gin.Context)
	UpdatePassword(c *gin.Context)
	GetUserProfile(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	DeleteUserProfile(c *gin.Context)
	VerifyEmail(c *gin.Context)

	RegenerateAccessAndRefreshToken(c *gin.Context)
}
type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// @Summary Generate Reset Password Token
// @Description Generate Reset Password Token
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {string} string "Reset password token generated successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/generateResetPasswordToken [post]
func (u *userController) GenerateResetPasswordToken(c *gin.Context) {
	var reqData struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	msg, err := u.userService.GenerateResetPasswordToken(reqData.Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Reset password token generated successfully", msg)
}

// @Summary Validate Reset Password Token
// @Description Validate Reset Password Token
// @Tags User
// @Accept json
// @Produce json
// @Param token query string true "Token"
// @Success 200 {string} string "Reset password token verified successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/validateResetPasswordToken [get]
func (u *userController) ValidateResetPasswordToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Token is required for reset password verification", nil)
		return
	}
	msg, err := u.userService.ValidateResetPasswordToken(token)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Reset password token verified successfully", msg)

}

// @Summary Update Password
// @Description Update Password
// @Tags User
// @Accept json
// @Produce json
// @Param password body string true "Password"
// @Param email body string true "Email"
// @Success 200 {string} string "Password updated successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/updatePassword [patch]
func (u *userController) UpdatePassword(c *gin.Context) {
	var reqData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	msg, err := u.userService.UpdatePassword(reqData.Email, reqData.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Password updated successfully", msg)
}

// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserSwagger true "User Request"
// @Success 201 {object} models.SuccessResponseUserSwagger "User created successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/register [post]
func (u *userController) RegisterUser(c *gin.Context) {
	var requsetData models.User
	if err := c.ShouldBindJSON(&requsetData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	user, err := u.userService.RegisterUser(requsetData)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// @Summary Verify Email
// @Description Verify Email
// @Tags User
// @Accept json
// @Produce json
// @Param token query string true "Token"
// @Success 200 {string}  "Email verified successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/verifyEmail [get]
func (u *userController) VerifyEmail(c *gin.Context) {
	secret := config.GetConfig().JwtSecret
	if secret == "" {
		utils.ErrorResponse(c, http.StatusInternalServerError, "JWT_SECRET_KEY environment variable not set", nil)
		return
	}

	fmt.Println("Secret key for email verification", secret)
	tokenString := c.Query("token")
	if tokenString == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Token is required for email verification", nil)
		return
	}

	fmt.Println("Token received for verification", tokenString)

	Email, err := auth.VerifyEmailVerificationToken(tokenString, secret)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	err = u.userService.VarifyEmail(Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Email verified successfully", nil)
}

// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} services.UserLoginData "User logged in successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/login [post]
func (u *userController) LoginUser(c *gin.Context) {
	var reqData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	user, err := u.userService.LoginUser(reqData.Email, reqData.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	if user.AccessToken != "" {
		c.SetCookie("access_token", user.AccessToken, 60*60*24*1, "/", "", false, true)
		c.SetCookie("refresh_token", user.RefreshToken, 60*60*24*7, "/", "", false, true)
	}

	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", user)
}

// @Summary Regenerate Access and Refresh Token
// @Description Regenerate Access and Refresh Token
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param token query string true "Token"
// @Success 200 {object} utils.AceesTokenAndRefreshToken  "Access and Refresh Token regenerated successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/regenerateAccessAndRefreshToken [get]
func (u *userController) RegenerateAccessAndRefreshToken(c *gin.Context) {
	var token string
	if cookieToken, err := c.Cookie("refresh_token"); err == nil {
		token = cookieToken
	}

	if token == "" {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			token = strings.TrimPrefix(strings.TrimSpace(authHeader), "Bearer ")
		}
	}

	if token == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Token not provided in Authorization header or cookie", nil)
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access", nil)
		return
	}

	log.Println("Claims:", claims)

	userID, ok := claims["id"].(string)

	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token: missing or invalid id", nil)
		return
	}
	userEmail, ok := claims["email"].(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token: missing or invalid email", nil)
		return
	}

	user, tokens, err := u.userService.RegenerateAccessAndRefreshToken(token)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	if user.Email != userID && user.Email != userEmail {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access", nil)
		return
	}

	if tokens.AccessToken != "" {
		c.SetCookie("access_token", tokens.AccessToken, 60*60*24*2, "/", "", false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, 60*60*24*7, "/", "", false, true)
	}

	utils.SuccessResponse(c, http.StatusOK, "Access and Refresh Token regenerated successfully", tokens)

}

// @Summary Get User Profile
// @Description Get User Profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.SuccessResponseUserSwagger"User profile retrived successfully"
// @Failure 401 {object} models.ErrorResponseUserSwagger "Unauthorized User"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/getUserProfile [get]
func (u *userController) GetUserProfile(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized User", err.Error())
		return
	}
	user, err := u.userService.GetUserProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User profile retrived successfully", user)
}

// @Summary Update User Profile
// @Description Update User Profile
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserSwagger true "User Request"
// @Security ApiKeyAuth
// @Success 200 {object} models.SuccessResponseUserSwagger"User profile updated successfully"
// @Failure 400 {object} models.ErrorResponseUserSwagger "Invalid input fields or JSON format"
// @Failure 401 {object} models.ErrorResponseUserSwagger "Unauthorized User"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/updateUserProfile [patch]
func (u *userController) UpdateUserProfile(c *gin.Context) {
	var requsetData models.User
	if err := c.ShouldBindJSON(&requsetData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized User", err.Error())
		return
	}
	user, err := u.userService.UpdateUserProfile(userID, requsetData)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User profile updated successfully", user)

}

// @Summary Delete User Profile
// @Description Delete User Profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} string "User profile deleted successfully"
// @Failure 401 {object} models.ErrorResponseUserSwagger "Unauthorized User"
// @Failure 500 {object} models.ErrorResponseUserSwagger "Internal server error"
// @Router /userAuth/deleteUserProfile [delete]
func (u *userController) DeleteUserProfile(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized User", err.Error())
		return
	}
	msg, err := u.userService.DeleteUserProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User profile deleted successfully", msg)
}
