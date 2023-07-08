package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	Set(key string, value interface{}, exp time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Delete(key string) *redis.IntCmd
}

type redisService struct {
	rdb *redis.Client
}

func NewRedisService(rdb *redis.Client) RedisService {
	return &redisService{rdb: rdb}
}

func (s *redisService) Set(key string, value interface{}, exp time.Duration) *redis.StatusCmd {
	return s.rdb.Set(context.Background(), key, value, exp)
}

func (s *redisService) Get(key string) *redis.StringCmd {
	return s.rdb.Get(context.Background(), key)
}

func (s *redisService) Delete(key string) *redis.IntCmd {
	return s.rdb.Del(context.Background(), key)
}
