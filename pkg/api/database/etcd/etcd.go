package etcd

import (
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/clientv3"
	"sync"
	"time"
)

var (
	once          sync.Once
	client        *v3Client
	DefaultConfig = &Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: time.Second * 5}
)

type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
}

type v3Client struct {
	client *clientv3.Client
}

func GetClientOnce() *v3Client {
	once.Do(func() {
		var err error
		client.client, err = newClient(DefaultConfig)
		if err != nil {
			panic(err)
		}
	})
	return client
}

func newClient(config *Config) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func (c *v3Client) Close() error {
	return c.client.Close()
}

func (c *v3Client) Put(ctx context.Context, key, value string) (*clientv3.PutResponse, error) {
	return c.client.Put(ctx, key, value)
}

func (c *v3Client) Get(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return c.client.Get(ctx, key)
}
