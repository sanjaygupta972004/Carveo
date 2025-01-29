package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheckResponse struct {
	Message string      `json:"message" example:"Server is running and healthy"`
	Status  string      `json:"status" example:"healthy"`
	Code    int         `json:"code" example:"200"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// @Summary Health Check Endpoint
// @Description Provides a simple health check to verify the server's operational status
// @Tags Health
// @Produce json
// @Success 200 {object} HealthCheckResponse
// @Router /health [get]
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
		SetupUserRoutes(rg, db)
	}

}

func SetupDefaultRouter(router *gin.Engine, port string) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("Welcome to Carveo, Server is running on port : %s", port),
		})
	})
}
