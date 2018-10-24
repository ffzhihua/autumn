package redis

import (
	"autumn/tools/cfg"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var REDIS redis.Conn

func InitRedis() {
	REDIS = redis_connect("default")
}

func CloseRedis() {
	REDIS.Close()
}

func redis_connect(project string) redis.Conn {

	server := fmt.Sprintf("%s:%s",
		cfg.Get("redis", project+".host").String(),
		cfg.Get("redis", project+".port").String())
	var err error
	option := redis.DialPassword(cfg.Get("redis", project+".passwd").String())
	c, err := redis.Dial("tcp", server, option)
	if err != nil {
		log.Fatal("[GIN-MYSQL(" + project + ")] connect to redis error:" + err.Error())
	}

	log.Println("[GIN-Redis(" + project + ")] connected success")

	return c
}
