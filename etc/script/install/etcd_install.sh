#!/bin/bash

#参数说明
#name      节点名称
#data-dir      指定节点的数据存储目录
#listen-peer-urls      自己监听本地的url
#listen-client-urls    《对外》提供服务的地址：比如 http://ip:2379,http://127.0.0.1:2379 ，客户端会连接到这里和 etcd 交互
#initial-advertise-peer-urls   该节点同伴监听地址，这个值会告诉集群中其他节点，其他节点能访问的url
#initial-cluster   集群中所有节点的信息，格式为 node1=http://ip1:2380,node2=http://ip2:2380,… 。注意：这里的 node1 是节点的 --name 指定的名字；后面的 ip1:2380 是 --initial-advertise-peer-urls 指定的值
#initial-cluster-state     新建集群的时候，这个值为 new ；假如已经存在的集群，这个值为 existing
#initial-cluster-token     创建集群的 token，这个值每个集群保持唯一。这样的话，如果你要重新创建集群，即使配置和之前一样，也会再次生成新的集群和节点 uuid；否则会导致多个集群之间的冲突，造成未知的错误
#advertise-client-urls     对外公告的该节点客户端监听地址，这个值会告诉集群中其他节点

docker run --restart=always --name=etcd1 -d \
	-p 2379:2379 \
	-p 2380:2380 \
	-v /var/etcd:/var/etcd \
	-v /etc/localtime:/etc/localtime \
	quay.io/coreos/etcd:v3.5.0 \
	etcd --name etcd1 \
	--auto-compaction-retention=1 --max-request-bytes=33554432 --quota-backend-bytes=8589934592 \
	--data-dir=/var/etcd/etcd-data \
	--listen-client-urls http://0.0.0.0:2379 \
	--listen-peer-urls http://0.0.0.0:2380 \
	--initial-advertise-peer-urls http://192.168.164.132:2380 \
	--advertise-client-urls http://192.168.164.132:2379 \
	--initial-cluster-state new \
	--initial-cluster-token etcd-cluster \
	--initial-cluster "etcd1=http://192.168.164.132:2380,etcd2=http://192.168.164.133:2380,etcd3=http://192.168.164.134:2380"

docker run --restart=always --name=etcd2 -d \
	-p 2379:2379 \
	-p 2380:2380 \
	-v /var/etcd:/var/etcd \
	-v /etc/localtime:/etc/localtime \
	quay.io/coreos/etcd:v3.5.0 \
	etcd --name etcd2 \
	--auto-compaction-retention=1 --max-request-bytes=33554432 --quota-backend-bytes=8589934592 \
	--data-dir=/var/etcd/etcd-data \
	--listen-client-urls http://0.0.0.0:2379 \
	--listen-peer-urls http://0.0.0.0:2380 \
	--initial-advertise-peer-urls http://192.168.164.133:2380 \
	--advertise-client-urls http://192.168.164.133:2379 \
	--initial-cluster-state new \
	--initial-cluster-token etcd-cluster \
	--initial-cluster "etcd1=http://192.168.164.132:2380,etcd2=http://192.168.164.133:2380,etcd3=http://192.168.164.134:2380"

docker run --restart=always --name=etcd3 -d \
	-p 2379:2379 \
	-p 2380:2380 \
	-v /var/etcd:/var/etcd \
	-v /etc/localtime:/etc/localtime \
	quay.io/coreos/etcd:v3.5.0 \
	etcd --name etcd3 \
	--auto-compaction-retention=1 --max-request-bytes=33554432 --quota-backend-bytes=8589934592 \
	--data-dir=/var/etcd/etcd-data \
	--listen-client-urls http://0.0.0.0:2379 \
	--listen-peer-urls http://0.0.0.0:2380 \
	--initial-advertise-peer-urls http://192.168.164.134:2380 \
	--advertise-client-urls http://192.168.164.134:2379 \
	--initial-cluster-state new \
	--initial-cluster-token etcd-cluster \
	--initial-cluster "etcd1=http://192.168.164.132:2380,etcd2=http://192.168.164.133:2380,etcd3=http://192.168.164.134:2380"

#查看leader
#docker exec -it etcd etcdctl -w table --endpoints http://127.0.0.1:2379 endpoint status --cluster