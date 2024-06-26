#### 节点通信

和集群中的任意节点通信，通信节点负责转发请求，并收集数据返回给客户端

mkdir -p /data/es

chown -R 1000:1000 /data/es


```yaml

version:  '3.9'
services:
  es01:
    image: elasticsearch:8.1.0
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

  es02:
    image: elasticsearch:8.1.0
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

  es03:
    image: elasticsearch:8.1.0
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


```

接口查询

```shell
查看所有索引
curl -sSL localhost:9200/_cat/indices?v

查看集群状态信息
curl localhost:9200/_cluster/stats?pretty

查看所有分片
curl localhost:9200/_cat/shards

查看集群健康状态
curl -XGET http://localhost:9200/_cluster/health?pretty

修改最大分片数，单节点默认最大分片数为1000
persistent持久，transient临时重启会丢失
curl -XPUT 'http://localhost:9200/_cluster/settings' -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "cluster.max_shards_per_node": 3000
  }
}'
查看集群配置
curl -XGET http://localhost:9200/_cluster/settings?pretty



```