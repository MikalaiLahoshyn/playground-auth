package redis

import (
	configs "auth/config"
	"auth/models"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

// handleRedisError handles redis error and wraps it into corresponding entity error.
func handleRedisError(name string, err error) error {
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis error[%s]: %w: %v", name, models.ErrNotFound, err.Error())
	}

	return fmt.Errorf("redis error[%s]: %w: %v", name, models.ErrInternal, err.Error())
}

func OpenDB(cfg configs.RedisDatabase) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Check if the connection is successful
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return client, nil
}
