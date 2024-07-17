package redisCache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"reflect"
	"testing"
	"time"
)

var test_cache *cache

func init() {
	test_cache = NewCache().(*cache)
}

func Test_cache_Delete(t *testing.T) {

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  *cache
		args    args
		wantErr bool
	}{
		{
			name:   "DeleteTest",
			fields: test_cache,
			args: args{
				key: "test",
			},
			wantErr: false,
		},
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
		fields *cache
		args   args
		want   string
	}{
		{
			name:   "combineKeyTest",
			fields: test_cache,
			args: args{
				key: "test",
			},
			want: "test",
		},
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
