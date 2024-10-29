package redis

import (
	"admin-api/common/config"
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

// SetupRedisDb Initialize the Redis instance
func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Server,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
