# goStorage
简单的分布式对象存储系统, 使用原生golang, rabbitmq搭建

## 基本数据流程如下:

## 启动方式
1. 首先开启rabbitmq服务, 并添加两个exchange. 可参照 rabbitmq.sh 文件内容
2. 开启 dataServeice
3. 开启 apiService
4. 访问 apiService

## 项目结构
项目结构如下, 运行的服务都包含在serveice中
```bash
go-storage
├─LICENSE                 // LICENSE
├─README.md               // README.md
├─go.mod                  // go module go.mod
├─go.sum                  // go module go.sum
├─rabbitmq.sh             // start rabbitmq container
├─services.sh             // start services container
├─services                // all services
|    ├─dataService        // 1. data service
|    |      ├─main.go     // main.go
|    |      ├─objects     // handle object put and get
|    |      ├─locate      // handle locate object
|    |      ├─heartbeat   // send heartbeat
|    ├─apiService         // 2. apiService
|    |     ├─main.go      // main.go
|    |     ├─objects      // handle object put and get request
|    |     ├─locate       // handle object locate request
|    |     ├─heartbeat    // listen data serveice heartbeat
├─pkg                     // self package
|  ├─utils                // utils
|  ├─stream               // stream handelr
|  ├─rabbitmq             // rabbitmq handler baseed on amqp
|  ├─netutils             // network heloper
|  ├─logs                 // local logs handler
```

## 运行说明
### 1.容器环境变量
* LISTEN_ADDRESS: 服务监听本地的端口
* STORAGE_ROOT: 数据服务存储数据的目录, 需要包含一个objects子目录, 数据会存储到子目录中

### 2.link RabbitMQ说明
服务启动时需要使用容器链接指定rabbitmq, 因为服务运行时会根据环境变量链接到RabbitMQ的服务. 如果没有使用容器运行的RabbitMQ, 也可以直接将其作为环境变量来启动服务的容器.
例如你的RabbitMQ运行在: 10.0.0.1, 服务端口为: 5672, 可以直接将其作为环境变量加入到启动服务的参数中:
-e RABBITMQ_PORT_5672_TCP_ADDR=10.0.0.1
-e RABBITMQ_PORT_5672_TCP_PORT=5672

## FAQ
any question please contact me