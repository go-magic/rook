package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"testing"
	"time"
)

func etcdDemo() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.164.128:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//value := `[{"path":"c:/tmp/nginx.log","topic":"web.log"},{"path":"d:/xxx/redis.log","topic":"redis.log"}]`
	value := `[{"path":"c:/tmp/nginx.log","topic":"web.log"},{"path":"d:/xxx/redis.log","topic":"redis.log"},{"path":"d:/xxx/mysql.log","topic":"mysql.log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	//_, err = cli.Put(ctx, "baodelu", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/collect_config")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func TestName(t *testing.T) {
	DefaultConfig.Endpoints = []string{"192.168.164.128:2379"}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	watch := GetClientOnce().Watch(ctx, "asd")
	for v := range watch {
		for _, v := range v.Events {
			fmt.Println(v.Kv)
		}
	}
}
