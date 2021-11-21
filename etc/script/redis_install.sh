#!/bin/bash

#--appendonly yes  开启redis 持久化
# redis-server /etc/redis/redis.conf以配置文件启动redis，加载容器内的conf文件，最终找到的是挂载的目

curl http://download.redis.io/redis-stable/redis.conf -o redis.conf

mkdir /etc/redis -p

cp redis.conf /etc/redis -f

docker rm -f redis

docker run -p 6379:6379 --name=redis \
      -v /etc/redis/redis.conf:/etc/redis/redis.conf \
      -v /data/redis/data:/data -d redis redis-server \
      /etc/redis/redis.conf --appendonly yes --requirepass "123456"