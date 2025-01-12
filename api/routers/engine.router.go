package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEngineRouter(rg *gin.RouterGroup, db *gorm.DB) {
	// engineRepository := repositories.NewEngineRepository(db)
	// engineService := services.NewEngineService(engineRepository)
	// engineHandler := controllers.NewEngineController(engineService)

	router := rg.Group("/engines")

	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from engine router",
			})
		})
	}
}
