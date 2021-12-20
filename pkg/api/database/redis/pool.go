package redis

import (
	"context"
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

func GetRedisDb()redis.Conn  {
	return pool.Get()
}

func GetRedisDbWithContext(ctx context.Context)(redis.Conn,error)  {
	return pool.GetContext(ctx)
}

func AddRedisDB(conn redis.Conn){
	defer func() {
		_ = conn.Close()
	}()
}
