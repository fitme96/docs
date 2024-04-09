#### docker swarm install

开放端口

-   7946 TCP/UDP 用来实现container发现服务的端口
-   4789 UDP container入口网络端口

```bash
下载docker离线安装包: 
https://download.docker.com/linux/ubuntu/dists/focal/pool/stable/amd64/


# 初始化master
docker swarm init --advertise-addr 192.168.190.103

docker swarm leave -f
# 加入工作节点
docker swarm join --token SWMTKN-1-14m2ihzx3yylou20fahewdzh2x30kruoowot8ngidudgvmnegr-crgue7bevr12dwlkwmqp7m5iy 192.168.65.121:2377

创建自定义overlay网络
docker network create -d overlay --subnet=10.10.0.0/16 --gateway=10.10.0.254  --attachable=true sec-network

```


#### 常用指令
```bash
docker stack deploy -c es.yml elasticsearch
docker service ls
```


