package redisCache

import (
	"context"
	"github.com/Thenecromance/OurStories/server/Interface"
)

const (
	list_internal_ = "paw"
)

// compile time check whether the cache type implements the Interface.CacheSupportList interface
var _ Interface.CacheSupportList = &cache{}

func (c *cache) ListPush(key string, value interface{}) error {
	c.internal = list_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	_, err := c.cli.LPush(c.ctx, key, value).Result()
	return err
}

func (c *cache) ListPop(key string) (any, error) {
	c.internal = list_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.LPop(c.ctx, key).Result()

}

func (c *cache) ListRange(key string, start, stop int64) ([]string, error) {
	c.internal = list_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.LRange(c.ctx, key, start, stop).Result()
}

func (c *cache) ListLength(key string) (int64, error) {
	c.internal = list_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	// todo: this usage might not correct, please check the documentation
	result, err := c.cli.LLen(c.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
