package rds

import (
	"github.com/hhcool/gtls/log"
	"go.uber.org/zap"
	"time"

	"github.com/go-redis/redis/v8"
)

/**********************LOCK 操作***********************/

var unlockCh = make(chan struct{}, 0)

func (r *ClientStruct) Lock(key string, Wait bool) bool {
	for {
		lockSuccess, err := r.SetNX(key, NewGoroutineId(), time.Second*10).Result()
		if err == nil && lockSuccess {
			go r.watchDog(key)
			return true
		} else {
			if !Wait {
				return false
			}
			time.Sleep(time.Millisecond * 30)
		}
	}
}
func (r *ClientStruct) UnLock(key string) {
	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1]
		then 
			return redis.call('del', KEYS[1]) 
		else 
			return 0 
		end
	`)
	if result, err := r.RunScript(script, key, NewGoroutineId()).Result(); err != nil || result == int64(0) {
		log.Error("解锁失败", zap.Error(err))
	} else {
		unlockCh <- struct{}{}
	}
}
func (r *ClientStruct) watchDog(key string) {
	expTicker := time.NewTicker(time.Second * 8)
	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1]
		then 
		  return redis.call('expire', KEYS[1], ARGV[2]) 
		else
		  return 0 
		end
	`)
	for {
		select {
		case <-expTicker.C:
			if result, err := r.RunScript(script, key, NewGoroutineId(), 10).Result(); err != nil || result == int64(0) {
				log.Errorf("锁续期失败[%s]：%v", key, err)
				return
			}
		case <-unlockCh:
			return
		}
	}
}
