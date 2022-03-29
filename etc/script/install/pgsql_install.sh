#!/bin/bash

pgdata=/root/lihao/script/data

rm -rf $pgdata

mkdir -p $pgdata

docker run --name postgres -e POSTGRES_USER="lx" -e POSTGRES_DB="lx" -e POSTGRES_PASSWORD="lx" -p 5432:5432 -v $pgdata:/var/lib/postgresql/data -d postgres