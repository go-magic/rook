#!/bin/bash

docker run --name=etcd --restart=always --net=host -v /etc/ssl/certs/:/etc/ssl/certs/ -d quay.io/coreos/etcd:v3.0.0_rc.1