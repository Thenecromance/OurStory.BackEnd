package redisCache

import (
	"context"
	"github.com/Thenecromance/OurStories/server/Interface"
)

// compile time check whether the cache type implements the Interface.CacheSupportList interface
var _ Interface.CacheSupportList = &cache{}

func (c *cache) ListPush(key string, value interface{}) error {

	ctx := context.Background()
	_, err := c.cli.LPush(ctx, key, value).Result()
	return err
}

func (c *cache) ListPop(key string) (any, error) {
	ctx := context.Background()
	return c.cli.LPop(ctx, key).Result()

}

func (c *cache) ListRange(key string, start, stop int64) ([]any, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cache) ListLength(key string) (int64, error) {
	//TODO implement me
	panic("implement me")
}
