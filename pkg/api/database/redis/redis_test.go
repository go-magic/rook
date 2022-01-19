package redis

import "testing"

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
	v, err := GetString("mykey")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(v)
}
