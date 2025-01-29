package controllers

import (
	"carveo/models"
	"carveo/services"
	"carveo/utils"
	"carveo/utils/auth"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetUserProfile(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	DeleteUserProfile(c *gin.Context)
	GetAllUsers(c *gin.Context)
	VerifyEmail(c *gin.Context)
}
type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

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

func (u *userController) VerifyEmail(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET_KEY")
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

func (u *userController) GetAllUsers(c *gin.Context) {
	panic("unimplemented")
}

func (u *userController) GetUserProfile(c *gin.Context) {
	panic("unimplemented")
}

func (u *userController) UpdateUserProfile(c *gin.Context) {
	panic("unimplemented")
}

func (u *userController) DeleteUserProfile(c *gin.Context) {
	panic("unimplemented")
}
