#### docker swarm install

开放端口

-   7946 TCP/UDP 用来实现container发现服务的端口
-   4789 UDP container入口网络端口

下载docker离线安装包: 
https://download.docker.com/linux/ubuntu/dists/focal/pool/stable/amd64/

安装docker-compose
 curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
 初始化master
docker swarm init --advertise-addr 192.168.190.103
加入工作节点
创建自定义overlay网络
docker network create -d overlay --subnet=10.10.0.0/16 --gateway=10.10.0.254  --attachable=true sec-network
nsenter 调试docker容器
 nsenter -t `docker inspect --format "{{ .State.Pid }}" 166` -n /bin/bash

#### 常用指令

docker stack deploy -c es.yml elasticsearch
docker service ls

#### dockerd 设置代理

root@sec:~# mkdir -p /etc/systemd/system/docker.service.d/
root@sec:~# cat >> /etc/systemd/system/docker.service.d/https_proxy.conf << EOF
Environment="HTTP_PROXY=http://192.168.60.49:7890/"
Environment="HTTPS_PROXY=http://192.168.60.49:7890/"
Environment="NO_PROXY=localhost,127.0.0.1,.example.com"s
EOF
systemctl daemon-reload 
systemctl restart docker