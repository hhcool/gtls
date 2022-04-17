package rds

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hhcool/gtls/log"
)

type Option struct {
	MasterName string
	Host       []string
	Password   string
	Type       int // 1单体 2集群 3哨兵
}

type ClientStruct struct {
	RedisClient  *redis.Client
	RedisCluster *redis.ClusterClient
	IsCluster    bool
}

var Client = new(ClientStruct)

const (
	Nil = redis.Nil
)

func NewRedis(o *Option) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch o.Type {
	case 1:
		client := redis.NewClient(&redis.Options{
			Addr:     o.getOneHost(),
			Password: o.getPassword(),
			DB:       0,
		})
		if _, err := client.Ping(ctx).Result(); err != nil {
			panic(err)
		} else {
			Client.setClient(client)
			log.Infof("初始化缓存库 单体")
		}
	case 2:
		cluster := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    o.getAllHost(),
			Password: o.getPassword(),
		})
		if _, err := cluster.Ping(ctx).Result(); err != nil {
			panic(err)
		} else {
			Client.setCluster(cluster)
			Client.IsCluster = true
			log.Infof("初始化缓存库 集群")
		}
	case 3:
		client := redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    o.MasterName,
			SentinelAddrs: o.getAllHost(),
			Password:      o.getPassword(),
			DB:            0,
		})
		if _, err := client.Ping(ctx).Result(); err != nil {
			panic(err)
		} else {
			Client.setClient(client)
			log.Infof("初始化缓存库 哨兵")
		}
	default:
		if len(o.Host) > 1 {
			o.Type = 2
		} else {
			o.Type = 1
		}
		NewRedis(o)
	}
}

// NewGoroutineId
// @Auth: oak  2021-10-15 18:36:22
// @Description:  线程ID
// @return int64
func NewGoroutineId() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func (o *Option) getAllHost() []string {
	if len(o.Host) == 0 {
		panic("redis host is required")
	}
	return o.Host
}
func (o *Option) getOneHost() string {
	if len(o.Host) == 0 {
		panic("redis host is required")
	}
	return o.Host[0]
}
func (o *Option) getPassword() string {
	return o.Password
}
func (r *ClientStruct) setClient(c *redis.Client) {
	r.RedisClient = c
}
func (r *ClientStruct) setCluster(c *redis.ClusterClient) {
	r.RedisCluster = c
}

func (r *ClientStruct) Del(k string) *redis.IntCmd {
	if r.IsCluster {
		return r.RedisCluster.Del(context.Background(), k)
	}
	return r.RedisClient.Del(context.Background(), k)

}
func (r *ClientStruct) Expire(k string, t time.Duration) *redis.BoolCmd {
	if r.IsCluster {
		return r.RedisCluster.Expire(context.Background(), k, t)
	}
	return r.RedisClient.Expire(context.Background(), k, t)
}

func (r *ClientStruct) RunScript(s *redis.Script, key string, args ...interface{}) *redis.Cmd {
	if r.IsCluster {
		return s.Run(context.Background(), r.RedisCluster, []string{key}, args...)
	} else {
		return s.Run(context.Background(), r.RedisClient, []string{key}, args...)
	}
}
