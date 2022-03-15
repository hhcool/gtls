package rds

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// HSet
// @Auth: oak  2021-10-25 15:26:57
// @Description:  向名称为key的hash中添加元素field
// @receiver r
// @param key
// @param value
// @return *redis.IntCmd
func (r *ClientStruct) HSet(key string, value ...interface{}) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.HSet(context.Background(), key, value...)
	}
	return r.RedisClient.HSet(context.Background(), key, value...)
}

// HGet
// @Auth: oak  2021-10-25 15:26:49
// @Description:  返回名称为key的hash中field对应的value
// @receiver r
// @param key
// @param field
// @return *redis.StringCmd
func (r *ClientStruct) HGet(key string, field string) *redis.StringCmd {
	if r.IsCluster {
		return r.RedisCluster.HGet(context.Background(), key, field)
	}
	return r.RedisClient.HGet(context.Background(), key, field)
}

// HGetAll
// @Auth: oak  2021-10-25 15:31:12
// @Description:	返回名称为key的hash中所有的键（field）及其对应的value
// @receiver r
// @param key
// @return *redis.StringStringMapCmd
func (r *ClientStruct) HGetAll(key string) *redis.StringStringMapCmd {
	if r.IsCluster {
		return r.RedisCluster.HGetAll(context.Background(), key)
	}
	return r.RedisClient.HGetAll(context.Background(), key)
}

// HIncrBy
// @Auth: oak  2021-10-26 00:14:54
// @Description:	将名称为key的hash中field的value增加integer
// @receiver r
// @param key
// @param field
// @param integer
// @return *redis.IntCmd
func (r *ClientStruct) HIncrBy(key string, field string, integer int64) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.HIncrBy(context.Background(), key, field, integer)
	}
	return r.RedisClient.HIncrBy(context.Background(), key, field, integer)
}
