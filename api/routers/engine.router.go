package routers

import (
	"carveo/api/controllers"
	"carveo/repositories"
	"carveo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEngineRouter(rg *gin.RouterGroup, db *gorm.DB) {
	engineRepository := repositories.NewEngineRepository(db)
	engineService := services.NewEngineService(engineRepository)
	engineHandler := controllers.NewEngineController(engineService)

	router := rg.Group("/engines")

	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from engine router",
			})
		})

		router.GET("/getAllEngine", engineHandler.GetEngines)
		router.GET("/:getAllEngineByID", engineHandler.GetEngine)
		router.POST("/createEngine", engineHandler.CreateEngine)
		router.PUT("/updateEngine/:engineID", engineHandler.UpdateEngine)
		router.DELETE("/deleteEngine/:engineID", engineHandler.DeleteEngine)
	}
}
