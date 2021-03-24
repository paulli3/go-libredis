# go-libredis
内部使用的redis操作类

##使用方式
```golang
redis := libredis.New(&libredis.RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	})
 redis.Set("key","v")
 
 ...
```
