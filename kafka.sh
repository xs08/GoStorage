# zookeeper
docker run -d --name zookeeper -p 2181:2181 -t wurstmeister/zookeeper

# kafka
docker run  -d --name kafka \
    -p 9092:9092 \
    -e KAFKA_BROKER_ID=0 \
    -e KAFKA_ZOOKEEPER_CONNECT=172.16.65.243:2181 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://172.16.65.243:9092 \
    -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -t wurstmeister/kafka

# kafka cluster
docker run -d --name kafka1 \
    -p 9093:9093 \
    -e KAFKA_BROKER_ID=1 \
    -e KAFKA_ZOOKEEPER_CONNECT=<宿主机IP>:2181 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://<宿主机IP>:9093 \
    -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9093 -t wurstmeister/kafka