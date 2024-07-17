package redisCache

import (
	"context"
	"github.com/Thenecromance/OurStories/server/Interface"
)

const (
	hash_internal_ = "bananas"
)

var _ Interface.CacheSupportHash = &cache{}

func (c *cache) HashSet(key string, field string, value interface{}) error {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	_, err := c.cli.HSet(c.ctx, key, field, value).Result()
	return err
}

func (c *cache) HashGet(key string, field string) (any, error) {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HGet(c.ctx, key, field).Result()
}

func (c *cache) HashDel(key string, field string) error {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	_, err := c.cli.HDel(c.ctx, key, field).Result()
	return err
}

func (c *cache) HashExists(key string, field string) (bool, error) {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HExists(c.ctx, key, field).Result()
}

func (c *cache) HashKeys(key string) ([]string, error) {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HKeys(c.ctx, key).Result()
}

func (c *cache) HashValues(key string) ([]string, error) {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HVals(c.ctx, key).Result()
}

func (c *cache) HashLength(key string) (int64, error) {

	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HLen(c.ctx, key).Result()
}

func (c *cache) HashGetAll(key string) (map[string]string, error) {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HGetAll(c.ctx, key).Result()
}

func (c *cache) HashSetObject(key string, obj interface{}) error {

	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HSet(c.ctx, key, obj).Err()
}

func (c *cache) HashGetObject(key string, obj interface{}) error {
	c.internal = hash_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.HGetAll(c.ctx, key).Scan(obj)
}
