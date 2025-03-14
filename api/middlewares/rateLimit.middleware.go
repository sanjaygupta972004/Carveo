package middlewares

import (
	"carveo/db"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	maxTokens          = 10
	refillRate         = 1
	refillRateInterval = 1 * time.Second
)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		redisClient := db.GetRedisClient()
		if redisClient == nil {
			log.Fatal("Redis client is not initialized")
		}
		userID := c.ClientIP()

		key := fmt.Sprintf("rate_limit:%s", userID)
		now := time.Now().Unix()
		ctx := context.Background()

		err := redisClient.Watch(ctx, func(tx *redis.Tx) error {
			lastRefillTime, _ := tx.HGet(ctx, key, "lastRefillTime").Int64()
			tokens, _ := tx.HGet(ctx, key, "tokens").Int64()

			if lastRefillTime == 0 {
				lastRefillTime = now
				tokens = maxTokens
			}

			timeElapsed := now - lastRefillTime
			newTokens := min(tokens+(timeElapsed*refillRate), maxTokens)
			if newTokens < 1 {
				return fmt.Errorf("rate limit exceeded")
			}

			_, err := tx.TxPipelined(ctx, func(p redis.Pipeliner) error {
				p.HSet(ctx, key, "toekens", newTokens-1)
				p.HSet(ctx, key, "lastRefilltime", now)
				p.Expire(ctx, key, 60*time.Second)
				return nil
			})

			c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", maxTokens))
			c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", newTokens-1))
			c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", now+lastRefillTime))

			return err
		}, key)

		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
		}
		c.Next()
	}
}
