package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool //创建redis连接池

func InitRedisPool(addr string,maxIdle int,MaxActive int,IdleTimeout time.Duration)  {
	pool = &redis.Pool{
		MaxIdle:maxIdle,
		MaxActive:MaxActive,
		IdleTimeout:IdleTimeout,
		Dial: func() (redis.Conn ,error){
			return redis.Dial("tcp",addr)
		},
	}
}

func GetPool() * redis.Pool{
	return pool
}