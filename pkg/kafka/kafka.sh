# zookeeper
docker run -d --name zookeeper -p 2181:2181  wurstmeister/zookeeper

# kafka with container
docker run -d --name kafka \
    --link zookeeper:zookeeper \
    -p 9092:9092 \
    -e KAFKA_BROKER_ID=0 \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://zookeeper:9092 \
    -e KAFKA_LISTENERS=PLAINTEXT://localhost:9092 \
    # -e KAFKA_CREATE_TOPICS=topic001:2:1 \ # create topic
    -it wurstmeister/kafka

# create topic
# kafka command locate /opt/kafka/bin
# reference https://blog.csdn.net/belonghuang157405/article/details/82149257
cd /opt/kafka_2.11-2.0.0/bin/

# 1. create topic
./kafka-topics.sh --create \
    --zookeeper zookeeper:2181 \
    --replication-factor 1 \
    --partitions 8 \
    --topic test
# 2. producer data. start console, enter msg and enter to send
./kafka-console-producer.sh --broker-list localhost:9092 --topic test

# 3. consumer msg
./kafka-console-consumer.sh \
    --bootstrap-server localhost:9092 \
    --topic test \
    --from-beginning

# 4. kafka-manager
docker pull sheepkiller/kafka-manager
docker run -d --name kafka-manager \
    --link zookeeper:zookeeper \
    -p 9000:9000 \
    -e ZK_HOSTS=zookeeper:2181 \
    --net=host \
    -it sheepkiller/kafka-manager firewall-cmd 


## 其他操作
# 1. 查看 topic
./kafka-topics.sh --list --zookeeper zookeeper:2181

# 2. 查看指定 topic 消息
./kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic test
