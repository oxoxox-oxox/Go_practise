package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStorage{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) Save(shortCode, longURL string) error {
	return r.client.Set(r.ctx, shortCode, longURL, 0).Err()
}

func (r *RedisStorage) Load(shortCode string) (string, bool) {
	val, err := r.client.Get(r.ctx, shortCode).Result()
	if err != nil {
		return "", false
	}
	return val, true
}
