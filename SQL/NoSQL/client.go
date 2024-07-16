package NoSQL

import (
	"github.com/go-redis/redis/v8"
)

var (
	setting *RedisSetting
)

func NewRedis() *redis.Client {
	return redis.NewClient(LoadSetting().ToRedisOption())
}
