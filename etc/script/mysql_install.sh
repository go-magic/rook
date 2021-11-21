#!/bin/bash

#privileged 获取root权限
docker run --name=mysql --restart=always --privileged=true \
          -p 3306:3306 -v /opt/mysql/data:/var/lib/mysql \
          -e MYSQL_ROOT_PASSWORD="aU0FEQ^#ad1" -d mysql