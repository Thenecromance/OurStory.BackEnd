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

func NewRedisWithDb(db int) *redis.Client {
	opt := LoadSetting().ToRedisOption()
	opt.DB = db
	return redis.NewClient(opt)
}
