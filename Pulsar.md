### 部署Pulsar

##### 创建bind目录

mkdir -p /data/zk/log
mkdir -p /data/bookie/{journal,ledgers}
mkdir -p /data/etcd
bin/pulsar-admin brokers list pulsar-cluster

##### 创建各组件

-   yml见下文
```shell
docker stack deploy -c zk.yml zoo

docker run --rm --network sec-network hub.bugfeel.net:8443/ck-test/pulsar:2.9.1 \
bin/pulsar initialize-cluster-metadata \
--cluster sec-pulsar-cluster \
--zookeeper zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181 \
--configuration-store zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181 \
--web-service-url http://broker01:8080,broker02:8080,broker03:8080 \
--web-service-url-tls https://broker01:8443,broker02:8443,broker03:8443 \
--broker-service-url pulsar://broker01:6650,broker02:6650,broker03:6650 \
--broker-service-url-tls pulsar+ssl://broker01:6651,broker02:6651,broker03:6651

docker stack deploy -c bookie.yml bookie
docker stack deploy -c pulsar.yml pulsar

```



### 清理环境

for i in zoo_zk pulsar_broker bookie_bookie;do for j in {01..03}; do docker service rm $i$j; done done

### 生产/消费测试

启动终端

docker run --rm --network sec-network -ti hub.bugfeel.net:8443/ck-test/pulsar:2.9.1 bash

创建topics

bin/pulsar-admin --admin-url [http://broker01:8080](http://10.10.0.16:8080) topics create persistent://public/default/test

生产

bin/pulsar-client --url pulsar://broker01:6650 produce persistent://public/default/test -n 2 -m "hello pulsar"

消费

bin/pulsar-client --url pulsar://broker01:6650 consume persistent://public/default/test -n 100 --subscription-name test-1 --subscription-type Exclusive

```yaml
zk.yaml

version:  '3.9'
services:
  zk01:
    image: hub.bugfeel.net:8443/ck-test/zookeeper:3.8.0
    environment:
      ZOO_MY_ID: 1
      ZOO_SERVERS: server.1=zoo_zk01:2888:3888;2181 server.2=zoo_zk02:2888:3888;2181 server.3=zoo_zk03:2888:3888;2181
    volumes:
      - type: bind
        source: /data/zk
        target: /data
      - type: bind
        source: /data/zk/log
        target: /datalog
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.zk==01
  zk02:
    image: hub.bugfeel.net:8443/ck-test/zookeeper:3.8.0
    environment:
      ZOO_MY_ID: 2
      ZOO_SERVERS: server.1=zoo_zk01:2888:3888;2181 server.2=zoo_zk02:2888:3888;2181 server.3=zoo_zk03:2888:3888;2181
    volumes:
      - type: bind
        source: /data/zk
        target: /data
      - type: bind
        source: /data/zk/log
        target: /datalog
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.zk==02
  zk03:
    image: hub.bugfeel.net:8443/ck-test/zookeeper:3.8.0
    environment:
      ZOO_MY_ID: 3
      ZOO_SERVERS: server.1=zoo_zk01:2888:3888;2181 server.2=zoo_zk02:2888:3888;2181 server.3=zoo_zk03:2888:3888;2181
    volumes:
      - type: bind
        source: /data/zk
        target: /data
      - type: bind
        source: /data/zk/log
        target: /datalog
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.zk==03
networks:
  default:
    external:
      name: sec-network
```

```yaml
bookie.yaml

version:  '3.9'
services:
  bookie01:
    image: hub.bugfeel.net:8443/ck-test/bookkeeper:4.14.4
    environment:
      BK_zkServers: "zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181"
      BK_zkLedgersRootPath: "/ledgers"
      BK_metadataServiceUri: "zk://zoo_zk01:2181;zoo_zk02:2181;zoo_zk03:2181/ledgers"
      BK_httpServerEnabled: "true"
      BK_bookiePort: "3181"
      BK_httpServerPort: "8080"
    volumes:
      - type: bind
        source: /data/bookie/journal
        target: /data/bookkeeper/journal
      - type: bind
        source: /data/bookie/ledgers
        target: /data/bookkeeper/ledgers
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.bookie==01
  bookie02:
    image: hub.bugfeel.net:8443/ck-test/bookkeeper:4.14.4
    environment:
      BK_zkServers: "zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181"
      BK_zkLedgersRootPath: "/ledgers"
      BK_metadataServiceUri: "zk://zoo_zk01:2181;zoo_zk02:2181;zoo_zk03:2181/ledgers"
      BK_httpServerEnabled: "true"
      BK_bookiePort: "3181"
      BK_httpServerPort: "8080"
    volumes:
      - type: bind
        source: /data/bookie/journal
        target: /data/bookkeeper/journal
      - type: bind
        source: /data/bookie/ledgers
        target: /data/bookkeeper/ledgers
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.bookie==02
  bookie03:
    image: hub.bugfeel.net:8443/ck-test/bookkeeper:4.14.4
    environment:
      BK_zkServers: "zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181"
      BK_zkLedgersRootPath: "/ledgers"
      BK_metadataServiceUri: "zk://zoo_zk01:2181;zoo_zk02:2181;zoo_zk03:2181/ledgers"
      BK_httpServerEnabled: "true"
      BK_bookiePort: "3181"
      BK_httpServerPort: "8080"
    volumes:
      - type: bind
        source: /data/bookie/journal
        target: /data/bookkeeper/journal
      - type: bind
        source: /data/bookie/ledgers
        target: /data/bookkeeper/ledgers
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.bookie==03
networks:
  default:
    external:
      name: sec-network


```

```yaml
pulsar.yaml

version:  '3.9'
services:
  broker01:
    image: hub.bugfeel.net:8443/ck-test/pulsar:2.9.1
    environment:
      clusterName: sec-pulsar-cluster
      zookeeperServers: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      configurationStore: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      brokerServicePort: "6650"
      brokerServicePortTls: "6651"
      webServicePort: "8080"
      webServicePortTls: "8443"
    command: >
      bash -c "bin/apply-config-from-env.py conf/broker.conf &&  \
               bin/apply-config-from-env.py conf/pulsar_env.sh &&  \
               exec bin/pulsar broker"
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.broker==01
  broker02:
    image: hub.bugfeel.net:8443/ck-test/pulsar:2.9.1
    environment:
      clusterName: sec-pulsar-cluster
      zookeeperServers: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      configurationStore: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      brokerServicePort: "6650"
      brokerServicePortTls: "6651"
      webServicePort: "8080"
      webServicePortTls: "8443"
    command: >
      bash -c "bin/apply-config-from-env.py conf/broker.conf && \
               bin/apply-config-from-env.py conf/pulsar_env.sh && \
               exec bin/pulsar broker"
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.broker==02
  broker03:
    image: hub.bugfeel.net:8443/ck-test/pulsar:2.9.1
    environment:
      clusterName: sec-pulsar-cluster
      zookeeperServers: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      configurationStore: zoo_zk01:2181,zoo_zk02:2181,zoo_zk03:2181
      brokerServicePort: "6650"
      brokerServicePortTls: "6651"
      webServicePort: "8080"
      webServicePortTls: "8443"
    command: >
      bash -c "bin/apply-config-from-env.py conf/broker.conf && \
               bin/apply-config-from-env.py conf/pulsar_env.sh && \
               exec bin/pulsar broker"
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.broker==03
networks:
  default:
    external:
      name: sec-network

```
