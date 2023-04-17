### 集群安装方式

1.  静态集群安装
2.  动态发现(discovery 设置节点数)
3.  DNS

etcdctl --endpoints=yulei_etcd01:2379,yulei_etcd02:2379,yulei_etcd03:2379 endpoint status -w table

#### Docker Swarm 静态部署
```yaml
version:  '3.9'
services:
  etcd01:
    image: hub.bugfeel.net:8443/ck-test/etcd:v3.5.2
    environment:
      ETCD_NAME: "etcd01"
      ETCD_DATA_DIR: "/var/lib/etcd"
      ETCD_LISTEN_PEER_URLS: "http://0.0.0.0:2380"
      ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379,http://0.0.0.0:4001"
      ETCD_INITIAL_ADVERTISE_PEER_URLS: "http://etcd_etcd01:2380"
      ETCD_ADVERTISE_CLIENT_URLS: "http://etcd_etcd01:2379"
      ETCD_INITIAL_CLUSTER: "etcd01=http://etcd_etcd01:2380,etcd02=http://etcd_etcd02:2380,etcd03=http://etcd_etcd03:2380"
      ETCD_INITIAL_CLUSTER_STATE: "new"
      ETCD_INITIAL_CLUSTER_TOKEN: "sec-etcd-cluster"
    volumes:
      - type: bind
        source: /data/etcd
        target: /var/lib/etcd
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.etcd==01
  etcd02:
    image: hub.bugfeel.net:8443/ck-test/etcd:v3.5.2
    environment:
      ETCD_NAME: "etcd02"
      ETCD_DATA_DIR: "/var/lib/etcd"
      ETCD_LISTEN_PEER_URLS: "http://0.0.0.0:2380"
      ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379,http://0.0.0.0:4001"
      ETCD_INITIAL_ADVERTISE_PEER_URLS: "http://etcd_etcd02:2380"
      ETCD_ADVERTISE_CLIENT_URLS: "http://etcd_etcd02:2379"
      ETCD_INITIAL_CLUSTER: "etcd01=http://etcd_etcd01:2380,etcd02=http://etcd_etcd02:2380,etcd03=http://etcd_etcd03:2380"
      ETCD_INITIAL_CLUSTER_STATE: "new"
      ETCD_INITIAL_CLUSTER_TOKEN: "sec-etcd-cluster"
    volumes:
      - type: bind
        source: /data/etcd
        target: /var/lib/etcd
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.etcd==02
  etcd03:
    image: hub.bugfeel.net:8443/ck-test/etcd:v3.5.2
    environment:
      ETCD_NAME: "etcd03"
      ETCD_DATA_DIR: "/var/lib/etcd"
      ETCD_LISTEN_PEER_URLS: "http://0.0.0.0:2380"
      ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379,http://0.0.0.0:4001"
      ETCD_INITIAL_ADVERTISE_PEER_URLS: "http://etcd_etcd03:2380"
      ETCD_ADVERTISE_CLIENT_URLS: "http://etcd_etcd03:2379"
      ETCD_INITIAL_CLUSTER: "etcd01=http://etcd_etcd01:2380,etcd02=http://etcd_etcd02:2380,etcd03=http://etcd_etcd03:2380"
      ETCD_INITIAL_CLUSTER_STATE: "new"
      ETCD_INITIAL_CLUSTER_TOKEN: "sec-etcd-cluster"
    volumes:
      - type: bind
        source: /data/etcd
        target: /var/lib/etcd
    deploy:
      mode: replicated
      replicas: 1
      endpoint_mode: dnsrr
      placement:
        constraints:
          - node.labels.etcd==03
networks:
  default:
    external:
      name: sec-network

```


#### 节点故障恢复