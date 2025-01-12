package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	rg := router.Group("/api")
	{
		SetupCarRouter(rg, db)
		SetupEngineRouter(rg, db)
	}

}
