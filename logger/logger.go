package logger

import (
	//"carveo/config"
	"carveo/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to load configuration:", err)
		os.Exit(1)
	}

	configApp := config.GetConfig()

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})

	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFile := filepath.Join(logDir, "api.log")
	writer, err := rotatelogs.New(
		logFile+".%Y%m%d",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationSize(100*1024*1024),
	)
	if err != nil {
		log.Fatalf("Failed to initialize log file rotation: %v", err)
	}

	multiWriter := io.MultiWriter(writer, os.Stdout)
	log.SetOutput(multiWriter)

	if configApp.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(logrus.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(logrus.DebugLevel)
	}

	log.WithFields(logrus.Fields{
		"app_env": configApp.GinMode,
	}).Info("Logger initialized successfully")
}

func GetLogger() *logrus.Logger {
	return log
}

func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

func Error(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

func Debug(message string, fields logrus.Fields) {
	log.WithFields(fields).Debug(message)
}

func Warn(message string, fields logrus.Fields) {
	log.WithFields(fields).Warn(message)
}
