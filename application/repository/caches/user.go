package caches

import (
	"github.com/Thenecromance/OurStories/SQL/NoSQL"
	"github.com/go-redis/redis/v8"
)

type UserCache interface {
}

type userCacheImpl struct {
	cli *redis.Client
}

func NewUserCache() UserCache {
	return &userCacheImpl{
		cli: NoSQL.NewRedis(),
	}
}
