package main

import (
	"carveo/api/routers"
	"carveo/config"
	"carveo/db"
	"carveo/db/migration"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cnfg config.Config

func init() {
	var err error = nil
	cnfg, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Println("Starting application...!")
	var err error = nil

	// initialized DB
	dbCon := &db.DBConfig{
		DbUser:     cnfg.DbUser,
		DbHost:     cnfg.DbHost,
		DbPort:     cnfg.DbPort,
		DbPassword: cnfg.DbPassword,
		DbName:     cnfg.DbName,
		SSLMode:    cnfg.SSLMode,
	}

	err = db.ConnectDB(dbCon)
	if err != nil {
		log.Fatalf("Failed to connect postgres db : %v", err)
	}
	defer db.DisConnectDB()
	db := db.DB
	// Migrate models
	if err := migration.MigrateModels(db); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)

	}
	// initialized gin server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome to Carveo, Server is running on port : %s", cnfg.Port),
		})
	})
	// setup routers
	routers.SetupRouter(router, db)
	log.Printf("Server is starting on port : %s", cnfg.Port)
	if err := router.Run(fmt.Sprintf(":%s", cnfg.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
