#!/bin/bash

docker run -itd -p 2379:2379 --restart=always -v /tmp/etcd-data.tmp:/etcd-data \
   --name etcd quay.io/coreos/etcd:v3.5.1 /usr/local/bin/etcd  \
   --listen-client-urls http://0.0.0.0:2379 \
   --advertise-client-urls http://0.0.0.0:2379