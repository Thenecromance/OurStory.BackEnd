package NoSQL

import "github.com/redis/go-redis/v9"

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
