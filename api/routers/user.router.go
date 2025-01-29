package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	// userRepo := repositories.NewUserRepository(db)
	// userService := services.NewUserService(userRepo)
	// userController := controllers.NewUserController(userService)

	// rg.POST("/register", userController.RegisterUser)
	// rg.POST("/login", userController.LoginUser)
	// rg.GET("/profile/:id", userController.GetUserProfile)
	// rg.PUT("/profile/:id", userController.UpdateUserProfile)
	// rg.DELETE("/profile/:id", userController.DeleteUserProfile)
	// rg.GET("/users", userController.GetAllUsers)

}
