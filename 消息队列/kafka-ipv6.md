## 先决条件

-   docker容器开启ipv6

## docker compose

-   conf_yulei网络可参考host.yaml network 创建

version: "3"
networks:
  conf_yulei:
    external: true
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    environment:
      # 匿名登录--必须开启
      - ALLOW_ANONYMOUS_LOGIN=yes
    networks:
      - conf_yulei
    #volumes:
      #- ./zookeeper:/bitnami/zookeeper
  # 该镜像具体配置参考 https://github.com/bitnami/bitnami-docker-kafka/blob/master/README.md
  kafka:
    image: 'bitnami/kafka:2.8.0'
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092
      # 客户端访问地址，更换成自己的
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://fa10::30c:28ee:fec7:32b8:9092
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      # 允许使用PLAINTEXT协议(镜像中默认为关闭,需要手动开启)
      - ALLOW_PLAINTEXT_LISTENER=yes
      # 关闭自动创建 topic 功能
      - KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=false
      # 全局消息过期时间 6 小时(测试时可以设置短一点)
      - KAFKA_CFG_LOG_RETENTION_HOURS=6
    #volumes:
      #- ./kafka:/bitnami/kafka
    networks:
      - conf_yulei

    depends_on:
      - zookeeper

#### 创建Topic

kafka-topics.sh --create --bootstrap-server [fa10::30c:28ee:fec7:32b8]:9092 --replication-factor 1 --partitions 1 --topic test

#### 启动消费者

docker  exec -ti root-kafka-1 bash
kafka-console-consumer.sh --bootstrap-server [fa10::30c:28ee:fec7:32b8]:9092 --topic test --from-beginning

#### 启动生产者

docker  exec -ti root-kafka-1 bash
kafka-console-producer.sh --broker-list [fa1028ee:fec7:32b8]:9092 --topic test