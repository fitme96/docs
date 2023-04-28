-   初始化

1.  增加SNAT(用于docker容器访问外部ipv6)

ip6tables -t nat -A POSTROUTING -s fc00:17:1:1::/80 -j MASQUERADE

2.  开启ipv6转发

-   临时开启ipv6转发

sysctl net.ipv6.conf.default.forwarding=1
sysctl net.ipv6.conf.all.forwarding=1

-   永久开启ipv6 转发

sed '$ a net.ipv6.conf.default.forwarding=1' /etc/sysctl.conf -i
sed '$ a net.ipv6.conf.all.forwarding=1' /etc/sysctl.conf -i
sysctl -p

### docker compose

1.  自定义网桥开启ipv6，指定子网用于容器分配全局ipv6地址

networks:
  yulei:
    driver: bridge
    enable_ipv6: true
    ipam:
      driver: default
      config:
        - subnet: 217.0.0.0/24 
        - subnet: fc00:17:1:1::/80
          gateway: fc00:17:1:1::1

### docker swarm

1.  手动新建docker_gwbridge

docker network create --ipv6 --subnet 172.19.0.0/16 --gateway 172.19.0.1 --gateway fc00:17:1:1::1 --subnet fc00:17:1:1::/80 --opt com.docker.network.bridge.name=docker_gwbridge  --opt com.docker.network.bridge.enable_icc=false  --opt com.docker.network.bridge.enable_ip_masquerade=true docker_gwbridge

_**注意**_ 容器会自动获取fe80::1/64子网的链接本地地址，宿主机ipv6地址不能同为一个子网