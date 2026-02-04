package redis

import (
	"context"
	"time"
)

func (rc *RedisClient) Set(key, value string) error {
	return rc.client.Set(context.Background(), key, value, 0).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Del(key string) error {
	return rc.client.Del(context.Background(), key).Err()
}

func (rc *RedisClient) Exists(key string) (bool, error) {
	result, err := rc.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (rc *RedisClient) TTL(key string) (time.Duration, error) {
	return rc.client.TTL(context.Background(), key).Result()
}

func (rc *RedisClient) Expire(key string, ttl int) error {
	return rc.client.Expire(context.Background(), key, time.Duration(ttl)*time.Second).Err()
}

func (rc *RedisClient) Keys(pattern string) ([]string, error) {
	return rc.client.Keys(context.Background(), pattern).Result()
}

func (rc *RedisClient) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return rc.client.Scan(context.Background(), cursor, match, count).Result()
}

func (rc *RedisClient) CountKeys(pattern string) (int64, error) {
	keys, err := rc.client.Keys(context.Background(), pattern).Result()
	if err != nil {
		return 0, err
	}
	return int64(len(keys)), nil
}
