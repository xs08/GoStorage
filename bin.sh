#!/bin/bash
# go env docker
docker run -d --name goenv \
    -v /d/project:/data \
    --link rabbitmq:rabbitmq
    -p 8888:8888 \
    -e LISTEN_PORT=8888 \
    -e STORAGE_ROOT=/home/bigsomg/project/go-storage/tmpStore/tmp \
    -it golang bash

# run
LISTEN_ADDRESS=:8888 STORAGE_ROOT=/data/tmpStore go run .