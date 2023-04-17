#### 节点通信

和集群中的任意节点通信，通信节点负责转发请求，并收集数据返回给客户端

mkdir -p /data/es

chown -R 1000:1000 /data/es

curl -XGET '[http://localhost:9200/_cluster/health?pretty'](http://localhost:9200/_cluster/health?pretty')

```yaml
es.yaml

version:  '3.9'
services:
  es01:
    image: hub.bugfeel.net:8443/ck-test/elasticsearch:8.1.0
    environment:
      node.name: elasticsearch_es01
      cluster.name: sec-es-cluster
      discovery.seed_hosts: elasticsearch_es02,elasticsearch_es03
      cluster.initial_master_nodes: elasticsearch_es01,elasticsearch_es02,elasticsearch_es03
      xpack.security.enabled: "false"
      bootstrap.memory_lock: "true"
      ES_JAVA_OPTS: "-Xms512m -Xmx512m"
    volumes:
      - type: bind
        source: /data/es
        target: /usr/share/elasticsearch/data
    ulimits:
      memlock:
        soft: -1
        hard: -1
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.es==01
  es02:
    image: hub.bugfeel.net:8443/ck-test/elasticsearch:8.1.0
    environment:
      node.name: elasticsearch_es02
      cluster.name: sec-es-cluster
      discovery.seed_hosts: elasticsearch_es01,elasticsearch_es03
      cluster.initial_master_nodes: elasticsearch_es01,elasticsearch_es02,elasticsearch_es03
      xpack.security.enabled: "false"
      bootstrap.memory_lock: "true"
      ES_JAVA_OPTS: "-Xms512m -Xmx512m"
    volumes:
      - type: bind
        source: /data/es
        target: /usr/share/elasticsearch/data
    ulimits:
      memlock:
        soft: -1
        hard: -1
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.es==02
  es03:
    image: hub.bugfeel.net:8443/ck-test/elasticsearch:8.1.0
    environment:
      node.name: elasticsearch_es03
      cluster.name: sec-es-cluster
      discovery.seed_hosts: elasticsearch_es01,elasticsearch_es02
      cluster.initial_master_nodes: elasticsearch_es01,elasticsearch_es02,elasticsearch_es03
      xpack.security.enabled: "false"
      bootstrap.memory_lock: "true"
      ES_JAVA_OPTS: "-Xms512m -Xmx512m"
    volumes:
      - type: bind
        source: /data/es
        target: /usr/share/elasticsearch/data
    ulimits:
      memlock:
        soft: -1
        hard: -1
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.es==03
networks:
  default:
    external:
      name: sec-network


```
