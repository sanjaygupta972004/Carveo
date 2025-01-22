package routers

import (
	"carveo/api/controllers"
	"carveo/repositories"
	"carveo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCarRouter(rg *gin.RouterGroup, db *gorm.DB) {
	carRepository := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepository)
	carHandler := controllers.NewCarController(carService)

	router := rg.Group("/cars")

	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from car router",
			})
		},
		)

		router.POST("/createCar", carHandler.CreateCar)
		router.GET("/getAllCars", carHandler.GetAllCars)
		router.GET("/getCarByID/:carID", carHandler.GetCarByID)
		router.GET("/getCarByBrand/:brandName/:isEngine", carHandler.GetCarByBrand)
		router.PATCH("/updateCar/:carID", carHandler.UpdateCar)
		router.DELETE("/deleteCar/:carID", carHandler.DeleteCar)
	}

}
