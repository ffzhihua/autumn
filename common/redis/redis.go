package myredis

import (
	"arutam/tools/cfg"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var REDIS redis.Conn
var RedisPool *redis.Pool

func InitRedis() {
	REDIS = redis_connect("default")
	RedisPool = getRedisPool("default")
}

func CloseRedis() {
	REDIS.Close()
}

func redis_connect(project string) redis.Conn {

	server := fmt.Sprintf("%s:%s",
		cfg.Get("redis", project+".host").String(),
		cfg.Get("redis", project+".port").String())
	var err error
	option := redis.DialPassword(cfg.Get("redis", project+".password").String())
	c, err := redis.Dial("tcp", server, option)
	if err != nil {
		log.Fatal("[GIN-MYSQL(" + project + ")] connect to redis error:" + err.Error())
	}

	log.Println("[GIN-Redis(" + project + ")] connected success")

	return c
}


func getRedisPool(project string) *redis.Pool {
	server := fmt.Sprintf("%s:%s",
		cfg.Get("redis", project+".host").String(),
		cfg.Get("redis", project+".port").String())
	option := redis.DialPassword(cfg.Get("redis", project+".password").String())

	return &redis.Pool{
		MaxIdle: 200,
		//MaxActive:   0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, option)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
