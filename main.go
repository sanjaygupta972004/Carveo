package main

import (
	"carveo/api/middlewares"
	"carveo/api/routers"
	"carveo/config"
	"carveo/db"
	"carveo/db/migration"
	"carveo/logger"
	"log"
	"os"

	_ "carveo/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title		Carevo Car Management API Documentation
// @version		1.0.0
// @description	API documentation for Carevo, a car management system for tracking, servicing, and managing vehicles.
// @termsOfService	https://carveo.com/terms

// @host		https://carveo.site
// @BasePath	/api/v1

// @contact.name		Support Team
// @contact.url		https://carveo.com/support
// @contact.email		sanjaygupta7212@gmailcom

// @securityDefinitions.apikey	      CarveoAPIKey
// @in						header
// @name						Authorization
// @description				JWT token required to access endpoints. Add the token in the "Authorization" header in the format: Bearer <JWT_TOKEN>.

// @tag.name		Car Management
// @tag.description	Endpoints related to car operations.

func main() {

	logger.Info("Starting application...!", logrus.Fields{})
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		logger.Error("Failed to load configuration:", logrus.Fields{
			"details": err,
		})
		log.Fatal("Critical Error: Shutting down application due to config failure.")
		return
	}
	logger.Info("Config loaded successfully", logrus.Fields{})

	configApp := config.GetConfig()

	// Initialize DB
	dbCon := &db.DBConfig{
		DbUser:     configApp.DbUser,
		DbHost:     configApp.DbHost,
		DbPort:     configApp.DbPort,
		DbPassword: configApp.DbPassword,
		DbName:     configApp.DbName,
		SSLMode:    configApp.SSLMode,
	}

	if err := db.ConnectDB(dbCon); err != nil {
		logger.Error("Failed to connect to Postgres DB:", logrus.Fields{
			"details": err,
		})

		log.Fatal("Critical Error: Shutting down application due to database connection failure.")
		os.Exit(1)

	}
	defer db.DisConnectDB()

	// Ensure DB connection is valid
	if db.DB == nil {
		logger.Error("Database connection is nil", logrus.Fields{})
	}

	dbInstance := db.DB

	// connect redis db
	if err := db.ConnectRedisDB(); err != nil {
		logger.Error(err.Error(), logrus.Fields{
			"detail": " Error while connecting redis database",
		})
		log.Fatal(err.Error())
	}

	// Migrate models
	if err := migration.MigrateModels(dbInstance); err != nil {
		logger.Error("Failed to migrate models:", logrus.Fields{
			"details": err,
		})
		log.Fatal("Critical Error: Shutting down application due to migration failure.")
		os.Exit(1)
	}

	// Initialize Gin server
	router := gin.New()

	// Global Middleware
	router.Use(middlewares.CorsMiddleWare())
	router.Use(middlewares.RecoverMiddleware())
	router.Use(middlewares.LoggerMiddleware())
	router.Use(middlewares.RateLimitMiddleware())

	// Setup routers
	routers.SetupRouter(router, dbInstance)
	routers.SetupHealthCheckRouter(router)

	// Default route
	routers.SetupDefaultRouter(router, configApp.Port)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("Swagger handler initialized...", logrus.Fields{})

	// Validate port before starting the server
	if configApp.Port == "" {
		logger.Error("Server port is not configured", logrus.Fields{})

	}

	logger.Info("Starting Gin server on port: %s", logrus.Fields{
		"port": configApp.DbPort,
	})
	if err := router.Run(":" + configApp.Port); err != nil {
		logger.Info("Failed to start server:", logrus.Fields{
			"details": err,
		})
		log.Fatal("Critical Error: Shutting down application due to server start failure.")
		os.Exit(1)
	}
}
