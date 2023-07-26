#!/bin/bash

docker run -d --name zookeeper --restart=always --publish 2181:2181 \
        --volume /etc/localtime:/etc/localtime zookeeper:latest

docker run -d --name kafka --restart=always --publish 9092:9092  \
        --link zookeeper  --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
        --env KAFKA_ADVERTISED_HOST_NAME=192.168.209.129  \
        --env KAFKA_ADVERTISED_PORT=9092  \
        --volume /etc/localtime:/etc/localtime  \
        wurstmeister/kafka