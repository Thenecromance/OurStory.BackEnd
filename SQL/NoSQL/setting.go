package NoSQL

import (
	"github.com/redis/go-redis/v9"
	"time"

	Config "github.com/Thenecromance/OurStories/utility/config"
)

type RedisSetting struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `json:"network,omitempty" yaml:"network"`
	// host:port address.
	Addr string `json:"addr,omitempty" yaml:"addr"`
	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `json:"username,omitempty" yaml:"username"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `json:"password,omitempty" yaml:"password"`

	// Database to be selected after connecting to the server.
	DB int `json:"db,omitempty" yaml:"db"`

	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetries int `json:"max_retries,omitempty" yaml:"max_retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `json:"min_retry_backoff,omitempty" yaml:"min_retry_backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `json:"max_retry_backoff,omitempty" yaml:"max_retry_backoff"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `json:"dial_timeout,omitempty" yaml:"dial_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `json:"read_timeout,omitempty" yaml:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `json:"write_timeout,omitempty" yaml:"write_timeout"`

	// Type of connection pool.
	// true for FIFO pool, false for LIFO pool.
	// Note that fifo has higher overhead compared to lifo.
	PoolFIFO bool `json:"pool_fifo,omitempty" yaml:"pool_fifo"`
	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize int `json:"pool_size,omitempty" yaml:"pool_size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `json:"min_idle_conns,omitempty" yaml:"min_idle_conns"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `json:"pool_timeout,omitempty" yaml:"pool_timeout"`

	// Enables read only queries on slave nodes.
	readOnly bool `json:"read_only,omitempty" yaml:"read_only"`

	// TLS Config to use. When set TLS will be negotiated.
	//TLSConfig *tls.Config `json:"tls_config,omitempty" yaml:"tls_config"`

	// Limiter interface used to implemented circuit breaker or rate limiter.
	//Limiter Limiter
}

func (r *RedisSetting) ToRedisOption() *redis.Options {
	result := redis.Options{
		Network:         r.Network,
		Addr:            r.Addr,
		Username:        r.Username,
		Password:        r.Password,
		DB:              r.DB,
		MaxRetries:      r.MaxRetries,
		MinRetryBackoff: r.MinRetryBackoff,
		MaxRetryBackoff: r.MaxRetryBackoff,
		DialTimeout:     r.DialTimeout,
		ReadTimeout:     r.ReadTimeout,
		WriteTimeout:    r.WriteTimeout,
		PoolFIFO:        r.PoolFIFO,

		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
		PoolTimeout:  r.PoolTimeout,
		//TLSConfig:          r.TLSConfig,
	}

	return &result
}

func defaultSetting() *RedisSetting {
	return &RedisSetting{
		Network: "tcp",
		Addr:    "localhost:6379",
	}
}

func LoadSetting() *RedisSetting {
	setting = defaultSetting()
	Config.Instance().LoadToObject("redis", setting)
	return setting
}
