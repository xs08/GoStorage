#!/bin/bash
# data service 1
docker run -d --name dataService1 \
    --link rabbitmq:rabbitmq \
    -p 21234:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmp/1:/data \
    -e LISTEN_ADDRESS=:1234 \
    -e STORAGE_ROOT=/data \
    -it golang bash

docker run -d --name dataService2 \
    --link rabbitmq:rabbitmq \
    -p 21235:1235 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmp/2:/data \
    -e LISTEN_ADDRESS=:1235 \
    -e STORAGE_ROOT=/data \
    -it golang bash

# api service 1
docker run -d --name apiService1 \
    --link rabbitmq:rabbitmq \
    -p 31234:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -e LISTEN_ADDRESS=:1234 \
    -it golang bash

# run
LISTEN_ADDRESS=:8888 STORAGE_ROOT=/data/tmpStore go run .