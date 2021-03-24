package libredis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisConfig struct {
	Host string
	Port int
	Auth string
	MaxIdle int
	MaxActive int
}


type mRedis struct {
	config RedisConfig
	pool   *redis.Pool
}


func New(c *RedisConfig) *mRedis {
	p := &mRedis{config: *c}
	p.connect()
	return p
}

func (c *mRedis) connect(){
	 c.pool = &redis.Pool{
		Wait:        true,
		MaxIdle:     1024,
		MaxActive:   5120,
		IdleTimeout: time.Duration(100 * time.Millisecond),
		Dial: func() (redis.Conn, error) {
			redisClient, err := redis.Dial(
				"tcp", fmt.Sprintf("%s:%d",c.config.Host,c.config.Port), redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				panic(fmt.Sprintf("connect redis fail %s:%d",c.config.Host,c.config.Port))
				return nil, err
			}

			if c.config.Auth != "" {
				_, err = redisClient.Do("AUTH", c.config.Auth)
				if err != nil {
					redisClient.Close()
					return nil, err
				}
			}

			return redisClient, nil
		},
	}
}

func (c *mRedis) Do(commond string, param ...interface{}) (interface{}, error) {
	client := c.pool.Get()
	defer client.Close()
	return client.Do(commond, param...)
}



func (c *mRedis) Get(key string) (string, error) {
	return redis.String(c.Do("get", key))
}

func (c *mRedis) Set(key string, value interface{}) (string, error) {
	return redis.String(c.Do("set", key,value))
}

func (c *mRedis) Del(key string) (string, error) {
	return redis.String(c.Do("del", key))
}

func (c *mRedis) SetEx(key string, value interface{}, timeout int) (string, error) {
	return redis.String(c.Do("setex", key,timeout,value))
}

func (c *mRedis) SetNx(key string, value interface{}) (int64, error) {
	return redis.Int64(c.Do("setnx", key,value))
}

func (c *mRedis) SetTimeout(key string, time int) (int64, error) {
	return redis.Int64(c.Do("expire", key,time))
}
func (c *mRedis) Ttl(key string) (int64, error) {
	return redis.Int64(c.Do("ttl",key))
}
//返回队列长度
func (c *mRedis) LPush(key string, value interface{}) (int64, error) {
	return redis.Int64(c.Do("lpush",key,value))
}
func (c *mRedis) Lpop(key string) (string, error) {
	return redis.String(c.Do("lpop",key))
}

func (c *mRedis) Lrange(key string, start int, stop int) ([]string, error) {
	return redis.Strings(c.Do("lrange",key,start,stop))
}

func (c *mRedis) RPush(key string, value interface{}) (int64, error) {
	return redis.Int64(c.Do("rpush",key,value))
}

func (c *mRedis) Rpop(key string) (string, error) {
	return redis.String(c.Do("rpop",key))
}

func (c *mRedis) Rrange(key string, start int, stop int) ([]string, error) {
	return redis.Strings(c.Do("rrange",key,start,stop))
}

func (c *mRedis) Incr(key string) (int64, error) {
	return redis.Int64(c.Do("incr",key))
}
func (c *mRedis) Decr(key string) (int64, error) {
	return redis.Int64(c.Do("decr",key))
}

func (c *mRedis) Hset(key string,subkey string,value interface{}) (int64, error)  {
	return redis.Int64(c.Do("hset",key,subkey,value))
}

func (c *mRedis) Hget(key string,subkey string) (string, error)  {
	return redis.String(c.Do("hget",key,subkey))
}

func (c *mRedis) Hmget(key string,subkeys ...interface{}) ([]string, error) {
	return redis.Strings(c.Do("hmget", append([]interface{}{key},subkeys...)...))
}

func (c *mRedis) Hmset(key string,data map[interface{}]interface{}) (string, error)  {
	return redis.String(c.Do("hmset", redis.Args{}.Add(key).AddFlat(data)...))
}
func (c *mRedis) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(c.Do("hgetall",key))
}

func (c *mRedis) Keys(key string) ([]string, error) {
	return redis.Strings(c.Do("keys",key ))
}
func (c *mRedis) Hlen(key string) (int64, error) {
	return redis.Int64(c.Do("hlen",key))
}
func (c *mRedis) Hdel(key string,subkey string) (int64, error) {
	return redis.Int64(c.Do("hdel",key,subkey))
}

func (c *mRedis) Hkeys(key string) ([]string, error) {
	return redis.Strings(c.Do("hkeys",key ))
}
func (c *mRedis) Hvals(key string) ([]string, error) {
	return redis.Strings(c.Do("hvals",key ))
}
func (c *mRedis) Hexists(key ,subkey string) (int64, error)  {
	return redis.Int64(c.Do("hexists",key,subkey))
}

func (c *mRedis) Sadd(key string,values ...interface{}) (int64, error) {
	return redis.Int64(c.Do("sadd ", redis.Args{}.Add(key).AddFlat(values)...))
}

func (c *mRedis) Smembers(key string) ([]string, error) {
	return redis.Strings(c.Do("smembers", key))
}

func (c *mRedis) SisMember(key string,val interface{}) (bool, error)  {
	return redis.Bool(c.Do("sismember", key,val))
}


func (c *mRedis) Zadd(key string,data map[int]string) (int64, error) {
	return redis.Int64(c.Do("zadd", redis.Args{}.Add(key).AddFlat(data)...))
}

func (c *mRedis) Zcard(key string) (int64, error) {
	return redis.Int64(c.Do("zcard",key))
}

func (c *mRedis) Zrange(key string,start,end int) ([]string, error) {
	return redis.Strings(c.Do("zrange", key,start,end))
}

func (c *mRedis) Zrevrange(key string,start,end int) ([]string, error) {
	return redis.Strings(c.Do("zrevrange", key,start,end))
}

func (c *mRedis) ZrangeWithScores(key string,start,end int) ([]string, error) {
	return redis.Strings(c.Do("zrange", key,start,end,"withscores"))
}

func (c *mRedis) ZrevrangeWithScores(key string,start,end int) ([]string, error) {
	return redis.Strings(c.Do("zrevrange", key,start,end,"withscores"))
}
