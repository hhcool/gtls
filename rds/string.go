package rds

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (r *ClientStruct) Set(k string, v interface{}, t time.Duration) *redis.StatusCmd {
	if r.IsCluster {
		return r.RedisCluster.Set(context.Background(), k, v, t)
	}
	return r.RedisClient.Set(context.Background(), k, v, t)
}
func (r *ClientStruct) SetNX(k string, v interface{}, t time.Duration) *redis.BoolCmd {
	if r.IsCluster {
		return r.RedisCluster.SetNX(context.Background(), k, v, t)
	}
	return r.RedisClient.SetNX(context.Background(), k, v, t)
}
func (r *ClientStruct) Get(k string) *redis.StringCmd {
	if r.IsCluster {
		return r.RedisCluster.Get(context.Background(), k)
	}
	return r.RedisClient.Get(context.Background(), k)
}
