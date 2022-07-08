## 1 根据proto文件，生成代码
```
goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 2 根据api文件，生成代码

```
goctl api go -api *.api -dir ./  --style=goZero
```

## 3 连接数据库，生成model文件
```
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/product" -table="*"  -dir="./model" -c
```

## 4 环境
### 4.1 kafka
```
brew install kafka
 
brew services start zookeeper

brew services start kafka

// 创建topic
./kafka-topics -create --bootstrap-server localhost:9092  --replication-factor 1 --partitions 1 --topic test1


// 列出topic
./kafka-topics --list --bootstrap-server localhost:9092

// 删除topic
./kafka-topics --delete --bootstrap-server localhost:9092  --topic seckill-topic


// 生产者
./kafka-console-producer --broker-list 127.0.0.1:9092 --topic t1

// 消费者(没有消费组)
./kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic t1  --from-beginning

// 消费者（指定消费组）开启一个消费者，指定消费者组. 生产者生产一条信息，消费者接收到
./kafka-console-consumer --bootstrap-server 0.0.0.0:9092 --topic t1 --group t1group

// 查看消费者信息
./kafka-consumer-groups -bootstrap-server 0.0.0.0:9092 --list

./kafka-consumer-groups -bootstrap-server 0.0.0.0:9092 --describe --group t1group

## 5 秒杀
```
grpcurl -plaintext -d '{"user_id": 111, "product_id": 10}' 127.0.0.1:9889 seckill.Seckill.SeckillOrder
```