#!/bin/bash
# data service 1
docker run -d --name dataService1 \
    --link rabbitmq:rabbitmq \
    -p 21234:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmpStore/tmp/1:/data \
    -e LISTEN_ADDRESS=127.0.0.1:1234 \
    -e STORAGE_ROOT=/data \
    -it golang bash

# data service 2
docker run -d --name dataService2 \
    --link rabbitmq:rabbitmq \
    -p 21235:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmpStore/tmp/2:/data \
    -e LISTEN_ADDRESS=127.0.0.1:1234 \
    -e STORAGE_ROOT=/data \
    -it golang bash

# run
LISTEN_ADDRESS=:8888 STORAGE_ROOT=/data/tmpStore go run .