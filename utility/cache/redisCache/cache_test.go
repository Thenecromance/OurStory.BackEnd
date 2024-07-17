package redisCache

import (
	"context"
	"github.com/Thenecromance/OurStories/server/Interface"
	"reflect"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		want Interface.ICache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCacheWithDb(t *testing.T) {
	type args struct {
		db int
	}
	tests := []struct {
		name string
		args args
		want Interface.ICache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCacheWithDb(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCacheWithDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_Delete(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cache_Get(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			got, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_GetPrefix(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			if got := c.GetPrefix(); got != tt.want {
				t.Errorf("GetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_GetSufix(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			if got := c.GetSufix(); got != tt.want {
				t.Errorf("GetSufix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_Prefix(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		prefix_ string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			c.Prefix(tt.args.prefix_)
		})
	}
}

func Test_cache_Set(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		key    string
		value  interface{}
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			if err := c.Set(tt.args.key, tt.args.value, tt.args.expire); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cache_Suffix(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		suffix_ string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			c.Suffix(tt.args.suffix_)
		})
	}
}

func Test_cache_clearInternal(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			c.clearInternal()
		})
	}
}

func Test_cache_combineKey(t *testing.T) {
	type fields struct {
		cli      *redis.Client
		ctx      context.Context
		internal string
		_prefix  string
		suffix   string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				cli:      tt.fields.cli,
				ctx:      tt.fields.ctx,
				internal: tt.fields.internal,
				_prefix:  tt.fields._prefix,
				suffix:   tt.fields.suffix,
			}
			if got := c.combineKey(tt.args.key); got != tt.want {
				t.Errorf("combineKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
