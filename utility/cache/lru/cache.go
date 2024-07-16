package lru

import (
	"fmt"
	"strings"
	"time"

	"github.com/Thenecromance/OurStories/server/Interface"
)

type cacheImpl struct {
	cli *Cache

	_prefix string
	_sufix  string
}

// Delete implements Interface.ICache.
func (c *cacheImpl) Delete(key string) error {
	c.cli.Remove(c.combineKey(key))
	return nil
}

// Get implements Interface.ICache.
func (c *cacheImpl) Get(key string) (any, error) {
	obj, ok := c.cli.Get(c.combineKey(key))
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return obj, nil
}

// GetPrefix implements Interface.ICache.
func (c *cacheImpl) GetPrefix() string {
	return c._prefix
}

// GetSufix implements Interface.ICache.
func (c *cacheImpl) GetSufix() string {
	return c._sufix
}

// Prefix implements Interface.ICache.
func (c *cacheImpl) Prefix(prefix_ string) {
	c._prefix = prefix_
}

// Set implements Interface.ICache.
func (c *cacheImpl) Set(key string, value interface{}, expire time.Duration) error {
	c.cli.Add(c.combineKey(key), value, time.Now().Add(expire))
	return nil
}

// Sufix implements Interface.ICache.
func (c *cacheImpl) Sufix(sufix_ string) {
	c._sufix = sufix_
}

func (c *cacheImpl) combineKey(key string) string {
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
	return &cacheImpl{
		cli: New(0),
	}
}
