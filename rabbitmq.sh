# This file for create rabbitmq on docker. 
# If you already have rabbitmq, just ignore it

# 1. get image
docker pull rabbitmq:management
# 2. start 
docker run -d --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    -e RABBITMQ_DEFAULT_USER=admin \
    -e RABBITMQ_DEFAULT_PASS=admin \
    -it rabbitmq:management

# 3. create exchange
# Two options to create it: 1. use rabbitmqadmin 2. use web interface
# My choice the easy way. open web browser, open http://127.0.0.1:15672/
# 1. if need login: enter the username and password, all are the same is amdin
# 2. click Exchange tab
# 3. click Add new exchange
# our service need two exchange: apiServers, dataServers
# create apiServers exChange
# enter Name: apiServers
# choice Type: fanout
# the other options we can use default
# 4. click Add exchange. and now, the All Exchanges table will dispaly a exchange named apiServers
# 5. folloing 3~4 to add dataServers
