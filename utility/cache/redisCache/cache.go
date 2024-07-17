package redisCache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"

	"github.com/Thenecromance/OurStories/SQL/NoSQL"
	"github.com/Thenecromance/OurStories/server/Interface"
)

/*
future work:
	change to a new framework to speed up data processing, just simulate CPU structures
	- L1: use local memory to store data
	- L2: use redis or other cache to store hotspot data
	- L3: use mysql for data persistence

fuck this crazy idea....
at least 1 server , 1 redis server ,1 mysql ,1 data profiler server
tooooooooooooooooooo expensive
*/

type cache struct {
	cli *redis.Client
	ctx context.Context

	internal string
	_prefix  string
	suffix   string
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
	return c.suffix
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

// Suffix implements Interface.ICache.
func (c *cache) Suffix(suffix_ string) {
	c.suffix = suffix_
}

func (c *cache) combineKey(key string) string {
	var builder strings.Builder
	if c.internal != "" {
		builder.WriteString(c.internal)
		builder.WriteString(".")
	}

	if c._prefix != "" {
		builder.WriteString(c._prefix)
		builder.WriteString(".")
	}

	builder.WriteString(key)
	if c.suffix != "" {
		builder.WriteString(".")
		builder.WriteString(c.suffix)
	}

	return builder.String()
}

func (c *cache) clearInternal() {
	c.internal = ""
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
