package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() *RedisClient {
	config := DefaultConfig()
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config *Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		MaxConnAge:   config.MaxConnAge,
		PoolTimeout:  config.PoolTimeout,
		IdleTimeout:  config.IdleTimeout,
	})

	return &RedisClient{
		client: rdb,
	}
}

func (rc *RedisClient) Ping() (string, error) {
	return rc.client.Ping(context.Background()).Result()
}

func (rc *RedisClient) Close() error {
	return rc.client.Close()
}
