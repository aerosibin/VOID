package cache

import (
	"context"
	"fmt"
	"time"

	"VOID/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewCache(cfg *config.Config) *Cache {
	if cfg == nil || cfg.Redis.Addr == "" {
		return &Cache{RedisClient: nil, Ctx: context.Background()}
	}
	rdb := redis.NewClient(&redis.Options{Addr: cfg.Redis.Addr})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("warning: redis unreachable: %v\n", err)
		return &Cache{RedisClient: nil, Ctx: ctx}
	}
	return &Cache{RedisClient: rdb, Ctx: ctx}
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	if c.RedisClient == nil {
		return nil
	}
	return c.RedisClient.Set(c.Ctx, key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	if c.RedisClient == nil {
		return "", nil
	}
	return c.RedisClient.Get(c.Ctx, key).Result()
}
