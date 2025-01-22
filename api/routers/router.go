package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary      Health Check
// @Description  Check if the server is running and healthy.
// @Tags         Health Check
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
func SetupHealthCheckRouter(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is running and healthy",
			"status":  "healthy",
			"code":    200,
			"error":   nil,
			"data":    nil,
		})
	})

}

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	rg := router.Group("/api/v1")
	{
		SetupCarRouter(rg, db)
		SetupEngineRouter(rg, db)
	}

}

func SetupDefaultRouter(router *gin.Engine, port string) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("Welcome to Carveo, Server is running on port : %s", port),
		})
	})
}
