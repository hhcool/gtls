package rds

import (
	"context"

	"github.com/go-redis/redis/v8"
)

/**********************SET 操作***********************/

// SAdd
// @Auth: oak  2021-10-17 01:50:18
// @Description: 向名称为key的set中添加元素member
// @receiver r
// @param k
// @param v
// @return *redis.IntCmd
func (r *ClientStruct) SAdd(key string, member ...interface{}) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.SAdd(context.Background(), key, member...)
	}
	return r.RedisClient.SAdd(context.Background(), key, member...)
}

// SRem
// @Auth: oak  2021-10-17 01:51:26
// @Description: 删除名称为key的set中的元素member
// @receiver r
// @param key
// @param member
// @return *redis.IntCmd
func (r *ClientStruct) SRem(key string, member ...interface{}) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.SRem(context.Background(), key, member...)
	}
	return r.RedisClient.SRem(context.Background(), key, member...)
}

// SPop
// @Auth: oak  2021-10-17 01:53:52
// @Description: 随机返回并删除名称为key的set中一个元素
// @receiver r
// @param key
// @return *redis.StringCmd
func (r *ClientStruct) SPop(key string) *redis.StringCmd {
	if r.IsCluster {
		return r.RedisCluster.SPop(context.Background(), key)
	}
	return r.RedisClient.SPop(context.Background(), key)
}

// SMembers
// @Auth: oak  2021-10-17 01:56:00
// @Description: 返回名称为key的set的所有元素
// @receiver r
// @param key
// @return *redis.StringSliceCmd
func (r *ClientStruct) SMembers(key string) *redis.StringSliceCmd {
	if r.IsCluster {
		return r.RedisCluster.SMembers(context.Background(), key)
	}
	return r.RedisClient.SMembers(context.Background(), key)
}

// SDiffStore
// @Auth: oak  2021-10-26 00:51:52
// @Description:	求差集并将差集保存到newKey的集合
// @receiver r
// @param newKey
// @param key
// @return *redis.IntCmd
func (r *ClientStruct) SDiffStore(newKey string, key ...string) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.SDiffStore(context.Background(), newKey, key...)
	}
	return r.RedisClient.SDiffStore(context.Background(), newKey, key...)
}
