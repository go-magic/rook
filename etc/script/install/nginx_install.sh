#!/bin/bash

docker run --name=nginx --restart=always -d nginx

docker cp nginx:/etc/nginx/conf.d ./

mkdir /etc/nginx/conf.d -p

cp conf.d/* /etc/nginx/conf.d -f

docker rm -f nginx

docker run --name=nginx --restart=always -v /etc/nginx/conf.d:/etc/nginx/conf.d -d nginx