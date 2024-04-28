
## Docker compose部署

### zk+kafka
```yaml
version: "3"
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    environment:
      # 匿名登录--必须开启
      - ALLOW_ANONYMOUS_LOGIN=yes
    volumes:
      - zookeeper:/bitnami/zookeeper

  kafka:
    image: 'bitnami/kafka:2.8.0'
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092
      # 客户端访问地址，更换成自己的
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.16.100:9092
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      # 允许使用PLAINTEXT协议(镜像中默认为关闭,需要手动开启)
      - ALLOW_PLAINTEXT_LISTENER=yes
      # 关闭自动创建 topic 功能
      - KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=false
      # 全局消息过期时间 6 小时(测试时可以设置短一点)
      - KAFKA_CFG_LOG_RETENTION_HOURS=6
    volumes:
      - kafka:/bitnami/kafka
    depends_on:
      - zookeeper
volumes:
  zookeeper:
  kafka:
```
***注意: KAFKA_CFG_ADVERTISED_LISTENERS需要配置为对外连接地址，在k8s中设置为Service_name***

### Kafka3.7
```yaml
version: "2"

services:
  kafka:
    image: docker.io/bitnami/kafka:3.7
    ports:
      - "9092:9092"
    volumes:
      - "kafka_data:/bitnami"
    environment:
      # KRaft settings
      - KAFKA_CFG_NODE_ID=0
      - KAFKA_CFG_PROCESS_ROLES=controller,broker
      - KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka:9093
      # Listeners
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.16.100:9092
      - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT
      - KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
      - KAFKA_CFG_INTER_BROKER_LISTENER_NAME=PLAINTEXT
volumes:
  kafka_data:

```

## 使用
#### 创建Topic
```bash
kafka-topics.sh --create --bootstrap-server 192.168.16.100:9092 --replication-factor 1 --partitions 1 --topic test
```
#### 启动消费者

```bash
kafka-console-consumer.sh --bootstrap-server 192.168.16.100:9092 --topic test --from-beginning
```
#### 启动生产者
```bash
kafka-console-producer.sh --broker-list 192.168.16.100:9092 --topic test
```

## 可视化工具

[kafdrop](https://hub.docker.com/r/obsidiandynamics/kafdrop)
```bash
docker run -d --rm -p 9000:9000 \
    -e KAFKA_BROKERCONNECT=host:port,host:port \
    -e JVM_OPTS="-Xms32M -Xmx64M" \
    -e SERVER_SERVLET_CONTEXTPATH="/" \
    obsidiandynamics/kafdrop:latest
```