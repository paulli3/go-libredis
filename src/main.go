package main

import (
	"apksystem/src/libredis"
	"fmt"
)

func main(){
	redis := libredis.New(&libredis.RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	})

	//nx, err := redis.Smembers("cc")
	//fmt.Println(nx,err)
	//
	//member, err := redis.SisMember("cc", "c")
	//
	//fmt.Println(member,err)
	//zadd, err := redis.Zadd("dd", map[int]string{3: "zhangsan", 4: "lishi"})
	scores, err := redis.ZrangeWithScores("dd",0,1)
	fmt.Println(scores, err)
}
