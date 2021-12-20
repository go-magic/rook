package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

func (c *v3Client) Watch(ctx context.Context, key string, opt ...clientv3.OpOption) clientv3.WatchChan {
	return c.client.Watch(ctx, key, opt...)
}
