# 1. get image
docker pull rabbitmq:management
# 2. start 
docker run -d --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    -e RABBITMQ_DEFAULT_USER=admin \
    -e RABBITMQ_DEFAULT_PASS=admin \
    -it rabbitmq:management

# 使用RabbitMQAdmin管理工具创建exchange
# 1. 远程服务安装开启rabbitmqadmin
# sudo apt-get install rabbitmq-server
# 2. 启用管理插件
# sudo rabbitmq-plugins enable rabbitmq_managment
# 3. 下载管理插件到本地
# wget lcoalhost:15672/cli/rabbitmqadmin
# 4. 使用python3执行命令
# python3 rabbitmqadmin declare exchange name=exchangeName type=exchangeType

# our service needs following exchange
rabbitmqadmin declare exchange name=apiServers type=fanout
rabbitmqadmin declare exchange name=dataServers type=fanout