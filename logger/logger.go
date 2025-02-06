package logger

import (
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

	if gin.Mode() == gin.ReleaseMode {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}
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
