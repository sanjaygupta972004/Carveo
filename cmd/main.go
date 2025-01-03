package main

import (
	"carveo/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var config Config

func init() {
	var err error
	config, err = LoadConfig()
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
		DbUser:     config.DbUser,
		DbHost:     config.DbHost,
		DbPort:     config.DbPort,
		DbPassword: config.DbPassword,
		DbName:     config.DbName,
		SSLMode:    config.SSLMode,
		TimeZone:   config.TimeZone,
	}

	err = db.ConnectDB(dbCon)
	if err != nil {
		log.Fatalf("Failed to connect postgres db : %v", err)
	}
	defer db.DisConnectDB()

	// db := db.DB

	// initialized gin server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome to Carveo, Server is run on port : %s", config.Port),
		})
	})
	if err := router.Run(":%s", config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
