package config

import (
	"backend/global"
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func initRedis() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	ctx := context.Background()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RedisDB = RedisClient
	log.Println("Redis initialized successfully")
}