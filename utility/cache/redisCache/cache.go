package redisCache

import (
	"context"
	"strings"
	"time"

	"github.com/Thenecromance/OurStories/SQL/NoSQL"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	cli *redis.Client

	_prefix string
	_sufix  string
}

// Delete implements Interface.ICache.
func (c *cache) Delete(key string) error {
	key = c.combineKey(key)
	ctx := context.Background()
	cmd := c.cli.Del(ctx, key)
	return cmd.Err()
}

// Get implements Interface.ICache.
func (c *cache) Get(key string) (any, error) {
	ctx := context.Background()
	cmd := c.cli.Get(ctx, c.combineKey(key))
	return cmd.Result()
}

// GetPrefix implements Interface.ICache.
func (c *cache) GetPrefix() string {
	return c._prefix
}

// GetSufix implements Interface.ICache.
func (c *cache) GetSufix() string {
	return c._sufix
}

// Prefix implements Interface.ICache.
func (c *cache) Prefix(prefix_ string) {
	c._prefix = prefix_
}

// Set implements Interface.ICache.
func (c *cache) Set(key string, value interface{}, expire time.Duration) error {
	ctx := context.Background()
	cmd := c.cli.Set(ctx, c.combineKey(key), value, expire)
	return cmd.Err()
}

// Sufix implements Interface.ICache.
func (c *cache) Sufix(sufix_ string) {
	c._sufix = sufix_
}

func (c *cache) combineKey(key string) string {
	var builder strings.Builder
	if c._prefix != "" {
		builder.WriteString(c._prefix)
		builder.WriteString(".")
	}

	builder.WriteString(key)
	if c._sufix != "" {
		builder.WriteString(".")
		builder.WriteString(c._sufix)
	}

	return builder.String()
}

func NewCache() Interface.ICache {
	return &cache{
		cli: NoSQL.NewRedis(),
	}
}

func NewCacheWithDb(db int) Interface.ICache {
	return &cache{
		cli: NoSQL.NewRedisWithDb(db),
	}
}
