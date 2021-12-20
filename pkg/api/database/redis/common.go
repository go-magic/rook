package redis

import (
	"github.com/garyburd/redigo/redis"
)

/*
执行命令
*/
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return conn.Do(commandName, args...)
}

/*
检测key是否存在
*/
func Exist(key string) (bool, error) {
	return redis.Bool(conn.Do("EXISTS", key))
}

/*
获取字符串
*/
func GetString(key string) (string, error) {
	return redis.String(conn.Do("GET", key))
}

/*
获取Bool
*/
func GetBool(key string) (bool, error) {
	return redis.Bool(conn.Do("GET", key))
}

/*
获取int
*/
func GetInt(key string) (int, error) {
	return redis.Int(conn.Do("GET", key))
}

/*
获取int64
*/
func GetInt64(key string) (int64, error) {
	return redis.Int64(conn.Do("GET", key))
}

/*
获取float64
*/
func GetFloat64(key string) (float64, error) {
	return redis.Float64(conn.Do("GET", key))
}
