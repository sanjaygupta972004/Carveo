package routers

import (
	"carveo/api/controllers"
	"carveo/api/middlewares"
	"carveo/repositories"
	"carveo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	authMiddleware := middlewares.JWTVerifyForUser(db)

	rg = rg.Group("/userAuth")

	rg.POST("/register", userController.RegisterUser)
	rg.GET("/verifyEmail", userController.VerifyEmail)
	rg.GET("/regenerateAccessAndRefreshToken", userController.RegenerateAccessAndRefreshToken)
	rg.POST("/generateResetPasswordToken", userController.GenerateResetPasswordToken)
	rg.GET("/validateResetPasswordToken", userController.ValidateResetPasswordToken)
	rg.PATCH("/resetPassword", userController.UpdatePassword)
	rg.POST("/login", userController.LoginUser)
	rg.GET("/getUserProfile", authMiddleware, userController.GetUserProfile)
	rg.PATCH("/updateUserProfile", authMiddleware, userController.UpdateUserProfile)
	rg.DELETE("/deleteUserProfile", authMiddleware, userController.DeleteUserProfile)
}
