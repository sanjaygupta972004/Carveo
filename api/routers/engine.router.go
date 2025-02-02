package routers

import (
	"carveo/api/controllers"
	"carveo/api/middlewares"
	"carveo/repositories"
	"carveo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEngineRouter(rg *gin.RouterGroup, db *gorm.DB) {
	engineRepository := repositories.NewEngineRepository(db)
	engineService := services.NewEngineService(engineRepository)
	engineHandler := controllers.NewEngineController(engineService)

	authMiddleware := middlewares.JWTVerifyForUser(db)

	router := rg.Group("/engines")
	router.Use(authMiddleware)
	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from engine router",
			})
		})

		router.GET("/getAllEngines", engineHandler.GetEngines)
		router.GET("/:getEngineByID/:engineID", engineHandler.GetEngine)
		router.POST("/createEngine/:carID", engineHandler.CreateEngine)
		router.PATCH("/updateEngine/:engineID", engineHandler.UpdateEngine)
		router.DELETE("/deleteEngine/:engineID", engineHandler.DeleteEngine)
	}
}
