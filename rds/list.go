package rds

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RPush
// @Auth: oak  2021-10-25 17:25:59
// @Description:  在名称为key的list尾添加一个值为value的元素
// @receiver r
// @param key
// @param value
// @return *redis.IntCmd
func (r *ClientStruct) RPush(key string, value ...interface{}) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.RPush(context.Background(), key, value...)
	}
	return r.RedisClient.RPush(context.Background(), key, value...)
}
