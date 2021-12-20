package redis


type Set struct {

}

/*
执行命令
*/
func SetWithTimeout(key string,timeout int,value string) (reply interface{}, err error) {
	return conn.Do("SETEX",key,timeout, value)
}
