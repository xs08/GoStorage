# This file for create service container

# 1. create data service with name dataService1
docker run -d --name dataService1 \
    --link rabbitmq:rabbitmq \
    -p 21234:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmp/1:/data \
    -e LISTEN_ADDRESS=:1234 \
    -e STORAGE_ROOT=/data \
    -it golang bash

# 2. create data service with name dataService2
docker run -d --name dataService2 \
    --link rabbitmq:rabbitmq \
    -p 21235:1235 \
    -v /home/bigsomg/project/go-storage:/app \
    -v /home/bigsomg/project/go-storage/tmp/2:/data \
    -e LISTEN_ADDRESS=:1235 \
    -e STORAGE_ROOT=/data \
    -it golang bash

# 3. create apiService with name apiService1
docker run -d --name apiService1 \
    --link rabbitmq:rabbitmq \
    -p 31234:1234 \
    -v /home/bigsomg/project/go-storage:/app \
    -e LISTEN_ADDRESS=:1234 \
    -it golang bash