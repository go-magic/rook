package redis

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool //创建redis连接池

func InitRedisPool(addr string, maxIdle int,
	MaxActive int, IdleTimeout time.Duration,
	passwd string) error {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialPassword(passwd))
		},
	}
	return nil
}

func GetPool() *redis.Pool {
	return pool
}

func GetRedisDb() (redis.Conn, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	return GetRedisDbWithContext(ctx)
}

func GetRedisDbWithContext(ctx context.Context) (redis.Conn, error) {
	return pool.GetContext(ctx)
}

func AddRedisDB(conn redis.Conn) {
	defer func() {
		_ = conn.Close()
	}()
}
