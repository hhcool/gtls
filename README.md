### golang 自用工具库

#### log 日志库 
> - 实现日志滚动落盘功能

- 使用方式
```
package main
import "github.com/hhcool/gtls/log"

func main() {
    // 同步日志到其他输出，如文件
    log.EnableSync(log.Option{
        ServerId string // 服务ID
        Format   Format // 日志类型
        Path     string // 日志文件路径，如果不传，不会存文件
        MaxAge   int    // 日志文件的生命周期，单位天
    })
}

// 获取一个安全的io.Write
// o.Write = log.SafeWriter()

// gorm日志
gorm.Config{
    Logger: log.NewGormLogger(),
}
```


#### mth 小数计算

#### tm 时间快捷方法


#### structs 针对struct的工具包
- 基于github.com/fatih/structs的分支包，感谢原作者
- 默认tag改为json，方便互转
- 增加了几个方法，比较零碎，自己看源码吧

#### 工具包
``` 
// 根据tag=default初始化struct
NewStructWithDefault(bean interface{})

// MD5
Md5(v string)
```


#### rds Redis库的封装
> 给redis套了个壳，主要是实现同一份代码同时兼容单体Redis和集群

##### string
- [x] set(key, value)             给数据库中名称为key的string赋予值value
- [x] get(key)                    返回数据库中名称为key的string的value
- [ ] getset(key, value)          给名称为key的string赋予上一次的value
- [ ] mget(key1, key2,…, key N)   返回库中多个string的value
- [x] setnx(key, value)           添加string，名称为key，值为value
- [ ] setex(key, time, value)     向库中添加string，设定过期时间time
- [ ] mset(key N, value N)        批量设置多个string的值
- [ ] msetnx(key N, value N)      如果所有名称为key i的string都不存在
- [ ] incr(key)                   名称为key的string增1操作
- [ ] incrby(key, integer)        名称为key的string增加integer
- [ ] decr(key)                   名称为key的string减1操作
- [ ] decrby(key, integer)        名称为key的string减少integer
- [ ] append(key, value)          名称为key的string的值附加value
- [ ] substr(key, start, end)     返回名称为key的string的value的子串

##### list
- [x] rpush(key, value)           在名称为key的list尾添加一个值为value的元素
- [ ] lpush(key, value)           在名称为key的list头添加一个值为value的 元素
- [ ] llen(key)                   返回名称为key的list的长度
- [ ] lrange(key, start, end)     返回名称为key的list中start至end之间的元素
- [ ] ltrim(key, start, end)      截取名称为key的list
- [ ] lindex(key, index)          返回名称为key的list中index位置的元素
- [ ] lset(key, index, value)     给名称为key的list中index位置的元素赋值
- [ ] lrem(key, count, value)     删除count个key的list中值为value的元素
- [ ] lpop(key)                   返回并删除名称为key的list中的首元素
- [ ] rpop(key)                   返回并删除名称为key的list中的尾元素
- [ ] blpop(key1, key2,… key N, timeout)  lpop命令的block版本。
- [ ] brpop(key1, key2,… key N, timeout)  rpop的block版本。
- [ ] rpoplpush(srckey, dstkey)           返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部

##### set
- [x] sadd(key, member)               向名称为key的set中添加元素member
- [x] srem(key, member)               删除名称为key的set中的元素member
- [x] spop(key)                       随机返回并删除名称为key的set中一个元素
- [ ] smove(srckey, dstkey, member)   移到集合元素
- [ ] scard(key)                      返回名称为key的set的基数
- [ ] sismember(key, member)          member是否是名称为key的set的元素
- [ ] sinter(key1, key2,…key N)       求交集
- [ ] sinterstore(dstkey, (keys))     求交集并将交集保存到dstkey的集合
- [ ] sunion(key1, (keys))            求并集
- [ ] sunionstore(dstkey, (keys))     求并集并将并集保存到dstkey的集合
- [ ] sdiff(key1, (keys))             求差集
- [x] sdiffstore(dstkey, (keys))      求差集并将差集保存到dstkey的集合
- [x] smembers(key)                   返回名称为key的set的所有元素
- [x] srandmember(key)                随机返回名称为key的set的一个元素

##### hash
- [x] hset(key, field, value)      向名称为key的hash中添加元素field
- [x] hget(key, field)             返回名称为key的hash中field对应的value
- [x] hincrby(key, field, integer) 将名称为key的hash中field的value增加integer
- [ ] hexists(key, field)          名称为key的hash中是否存在键为field的域
- [ ] hdel(key, field)             删除名称为key的hash中键为field的域
- [ ] hlen(key)                    返回名称为key的hash中元素个数
- [ ] hkeys(key)                   返回名称为key的hash中所有键
- [ ] hvals(key)                   返回名称为key的hash中所有键对应的value
- [x] hgetall(key)                 返回名称为key的hash中所有的键（field）及其对应的value

##### 锁
```
Lock(key string, Wait bool)
UnLock(key)
```