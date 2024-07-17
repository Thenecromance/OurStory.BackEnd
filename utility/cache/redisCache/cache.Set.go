package redisCache

import (
	"context"
	"github.com/Thenecromance/OurStories/server/Interface"
)

var _ Interface.CacheSupportSet = &cache{}

const (
	set_internal_ = "set"
)

func (c *cache) SetAdd(key string, value interface{}) error {
	c.internal = set_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	_, err := c.cli.SAdd(c.ctx, key, value).Result()
	return err
}

func (c *cache) SetRemove(key string, value interface{}) error {
	c.internal = set_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	_, err := c.cli.SRem(c.ctx, key, value).Result()
	return err

}

func (c *cache) SetIsMember(key string, value interface{}) (bool, error) {
	c.internal = set_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.SIsMember(c.ctx, key, value).Result()
}

func (c *cache) SetMembers(key string) ([]string, error) {
	c.internal = set_internal_
	defer c.clearInternal()
	key = c.combineKey(key)
	c.ctx = context.Background()

	return c.cli.SMembers(c.ctx, key).Result()
}
