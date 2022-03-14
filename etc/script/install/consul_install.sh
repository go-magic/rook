#!/bin/bash

# 安装consul集群
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server1 consul:latest agent \
	-server -bootstrap -ui -node=server1 -client='0.0.0.0' \
	-node='192.168.164.132' -bind='0.0.0.0' -advertise='192.168.164.132'

docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server2 consul:latest agent \
	-server -bootstrap -ui -node=server2 -client='0.0.0.0' \
	-node='192.168.164.133' -bind='0.0.0.0' -advertise='192.168.164.133' \
	 -join=192.168.164.132

docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server3 consul:latest agent \
	-server -bootstrap -ui -node=server3 -client='0.0.0.0' \
	-node='192.168.164.134' -bind='0.0.0.0' -advertise='192.168.164.134' \
	 -join=192.168.164.132

