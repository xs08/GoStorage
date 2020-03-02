# goStorage
简单的分布式对象存储系统, 使用原生golang, rabbitmq搭建

## 基本数据流程如下:

## 启动方式
1. 首先开启rabbitmq服务, 并添加两个exchange. 可参照 rabbitmq.sh 文件内容
2. 开启 dataServeice
3. 开启 apiService
4. 访问 apiService