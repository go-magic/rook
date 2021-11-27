package redis

import (
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

/*
初始化redis链接
*/
func InitRedis(addr string, passwd string) error {
	var err error
	conn, err = redis.Dial("tcp", addr)
	if err != nil {
		return err
	}
	_, err = conn.Do("auth", passwd)
	return err
}

/*
关闭redis链接
*/
func Close() error {
	return conn.Close()
}


type Reply interface {
	RedisSuccess()bool
}
