#!/bin/bash

#参数说明
#-client 0.0.0.0表示其他机器服务可以访问，127.0.0.1表示只能本机服务可访问
#-bind 0.0.0.0表示其他机器服务可以访问，127.0.0.1表示只能本机服务可访问
#-node集群中的ip地址
#-advertise对外暴露的ip地址
# 安装consul集群
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server1 consul:1.13.1 agent \
	-server -bootstrap-expect=3 -ui -node=server1 -client='0.0.0.0' \
	-node='192.168.164.132' -bind='0.0.0.0' -advertise='192.168.164.132'

docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server2 consul:1.13.1 agent \
	-server -bootstrap-expect=3 -ui -node=server2 -client='0.0.0.0' \
	-node='192.168.164.133' -bind='0.0.0.0' -advertise='192.168.164.133' \
	 -join=192.168.164.132

docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul_server3 consul:1.13.1 agent \
	-server -bootstrap-expect=3 -ui -node=server3 -client='0.0.0.0' \
	-node='192.168.164.134' -bind='0.0.0.0' -advertise='192.168.164.134' \
	 -join=192.168.164.132

#单机版
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 \
	--restart=always --name=consul consul:1.13.1 agent \
	-server -bootstrap -ui -node=node1 -client='0.0.0.0'

