package cache

import (
	"context"
	"os"
	"time"

	"go-pos/internal/core/port"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewCache(ctx context.Context) (port.CacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(res), nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) DeleteByPrefix(ctx context.Context, prefix string) error {
	var (
		cursor uint64
		keys   []string
	)

	for {
		var err error
		keys, cursor, err = r.client.Scan(ctx, cursor, prefix, 100).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			if err := r.client.Del(ctx, key).Err(); err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
