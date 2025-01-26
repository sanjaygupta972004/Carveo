package main

import (
	"carveo/api/middlewares"
	"carveo/api/routers"
	"carveo/config"
	"carveo/db"
	"carveo/db/migration"
	"log"

	_ "carveo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title		Carevo Car Management API Documentation
// @version		1.0.0
// @description	API documentation for Carevo, a car management system for tracking, servicing, and managing vehicles.
// @termsOfService	https://carveo.com/terms

// @host		http://localhost:8090
// @BasePath	/api/v1

// @contact.name		Support Team
// @contact.url		https://carveo.com/support
// @contact.email		sanjaygupta07054@gmailcom

// @securityDefinitions.apikey	      CarveoAPIKey
// @in						header
// @name						Authorization
// @description				JWT token required to access endpoints. Add the token in the "Authorization" header in the format: Bearer <JWT_TOKEN>.

// @tag.name		Car Management
// @tag.description	Endpoints related to car operations.

var ConfigApp config.Config

func init() {
	var err error = nil
	ConfigApp, err = config.LoadConfig()
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
		DbUser:     ConfigApp.DbUser,
		DbHost:     ConfigApp.DbHost,
		DbPort:     ConfigApp.DbPort,
		DbPassword: ConfigApp.DbPassword,
		DbName:     ConfigApp.DbName,
		SSLMode:    ConfigApp.SSLMode,
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
	// global middleware
	router.Use(middlewares.CorsMiddleWare())
	router.Use(middlewares.RecoverMiddleware())
	// setup routers
	routers.SetupRouter(router, db)
	// swagger handler for gin
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// setup default router
	routers.SetupDefaultRouter(router, ConfigApp.Port)

	log.Println("Swagger handler initialized...")
	log.Println("Starting Gin server on port:", ConfigApp.Port)

	// setup health check router
	routers.SetupHealthCheckRouter(router)
	// run server
	router.Run(":" + ConfigApp.Port)

}
