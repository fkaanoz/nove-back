package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClientWithOptions(options *redis.Options) *Redis {
	return &Redis{
		Client: redis.NewClient(options),
	}
}

func NewRedisClientWithURL(url string) (*Redis, error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return &Redis{Client: redis.NewClient(options)}, nil
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil

}

func (r *Redis) Set(key string, value string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := r.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}

	return true
}
