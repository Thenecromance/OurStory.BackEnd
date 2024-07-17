package Interface

import "time"

// ICache is the interface for cache data in local machine which will be used to store the data temporarily
type ICache interface {
	Prefix(prefix_ string)
	GetPrefix() string

	Suffix(sufix_ string)
	GetSufix() string

	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}

type CacheSupportList interface {
	ListPush(key string, value interface{}) error
	ListPop(key string) (any, error)
	ListRange(key string, start, stop int64) ([]string, error)
	ListLength(key string) (int64, error)
}

type CacheSupportSet interface {
	SetAdd(key string, value interface{}) error
	SetRemove(key string, value interface{}) error
	SetIsMember(key string, value interface{}) (bool, error)
	SetMembers(key string) ([]string, error)
}

type CacheSupportHash interface {
	HashSet(key string, field string, value interface{}) error
	HashGet(key string, field string) (any, error)
	HashDel(key string, field string) error
	HashExists(key string, field string) (bool, error)
	HashKeys(key string) ([]string, error)
	HashValues(key string) ([]string, error)
	HashLength(key string) (int64, error)
	HashGetAll(key string) (map[string]string, error)

	HashSetObject(key string, obj interface{}) error
	HashGetObject(key string, obj interface{}) error
}
