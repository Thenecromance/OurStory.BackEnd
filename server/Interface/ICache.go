package Interface

import "time"

// ICache is the interface for cache data in local machine which will be used to store the data temporarily
type ICache interface {
	Prefix(prefix_ string)
	GetPrefix() string

	Sufix(sufix_ string)
	GetSufix() string

	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}
