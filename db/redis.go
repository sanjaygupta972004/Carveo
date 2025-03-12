package db

import (
	"carveo/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedisDB() error {
	config := config.GetConfig()
	opt, err := redis.ParseURL(config.RedisConfig.RedisURL)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis connection failed!")
		return err
	} else {
		fmt.Println("Redis connected successfully...")
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
