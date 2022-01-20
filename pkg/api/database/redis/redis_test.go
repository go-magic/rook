package redis

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func init() {
	if err := InitRedis("192.168.164.128:6379", "123456"); err != nil {
		panic(err)
	}
}

func TestRedis(t *testing.T) {
	s, err := Do("SET", "mykey", "superWang")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(s)
	v, err := GetString(conn, "mykey")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(v)
}

func TestInitRedisPool(t *testing.T) {
	InitRedisPool("192.168.164.128:6379",
		8, 0, 100, "123456")
	GetPool().Get().Do("setex", "abc", 10, "123")
	t.Log(redis.String(GetPool().Get().Do("get", "abc")))
}
