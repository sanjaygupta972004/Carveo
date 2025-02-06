package middlewares

import (
	"carveo/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestPath := c.Request.URL.Path
		queryParams := c.Request.URL.RawQuery

		c.Next()

		// Log request details
		logger.Info("Request Processed", logrus.Fields{
			"status":     c.Writer.Status(),
			"query":      queryParams,
			"method":     c.Request.Method,
			"path":       requestPath,
			"latency":    time.Since(startTime),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})

		// Log errors if any
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				fmt.Printf("error in logger middleware : %v", err.Error())
				logger.Error("Error occured in request process ", logrus.Fields{
					"status": c.Writer.Status(),
					"path":   requestPath,
				})
			}
		}
	}
}
