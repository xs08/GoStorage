# 1. get image
docker pull rabbitmq:management
# 2. start 
docker run -d --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    -e RABBITMQ_DEFAULT_USER=admin \
    -e RABBITMQ_DEFAULT_PASS=admin \
    -it rabbitmq:management
