## 先决条件

- master依赖brctl,执行apt install -y bridge-utils

## 背景
- docker overlay 网络模式下，主机网络无法与overlay网络直接通信，只能通过端口映射等方式实现访问，通过veth打通主机网络与overlay网络。

## 实施

### master节点操作
```bash
1.  docker netns 创建软链接, 通过ip netns 可以看到network namespace,docker默认放到/var/run/docker/netns下，ip netns不可见

rm -rf /var/run/netns
ln -sv /var/run/docker/netns /var/run/netns
ip netns #查看docker创建的ns

2.  添加veth pair

ip link add sec01 type veth peer name sec02

3.  确认overlay网络ns

ip netns exec 1-p20dp5pq1y ip a

4.  将veth pair 一端sec01放入overlay网络 ns

ip link set sec01 netns 1-p20dp5pq1y

5.  将sec01加入br0

ip netns exec 1-p20dp5pq1y brctl addif br0 sec01

6.  veth pair sec02主机一端增加ip地址并up

ip addr add 10.0.1.252/24 dev sec02
ip link set sec02 up

7.  up sec01一端

ip netns exec 1-p20dp5pq1y ip link set sec01 up

```
***至此master主机可以通过本机直接访问overlay网络容器，下面通过配置FDB表以及ARP记录实现跨主机访问overlay网络容器***


### node节点操作
```bash
1. 获取masterIP地址及veth pair sec02的macaddr

2. ip netns exec 1-p20dp5pq1y bash

3. bridge fdb append $veth_sec01_macaddr dev vxlan0 dst $master_ip

4. ip neighbor add $sec02_ipaddr lladdr $veth_sec02_macaddr dev vxlan0

```
