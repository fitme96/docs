##先决条件

-   master依赖brctl,执行apt install -y bridge-utils

## 实施

### master节点操作
```bash
1.  docker netns创建软链接,通过ip netns 可以看到network namespace,docker默认放到/var/run/docker/netns下，ip netns不可见

   rm -rf /var/run/netns
   ln -sv /var/run/docker/netns /var/run/netns
   ip netns #查看docker创建的ns

2.  添加veth pair

   ip link add sec01 type veth peer name sec02

3.  确认br0 ns

   ip netns exec 1-p20dp5pq1y ip a

4.  将veth pair 一端sec01放入br0 ns

   ip link set sec01 netns 1-p20dp5pq1y

5.  将sec01加入br0

   ip netns exec 1-p20dp5pq1y brctl addif br0 sec01

6.  veth pair sec02主机一端增加ip地址并up

   ip addr add 10.0.1.252/24 dev sec02
   ip link set sec02 up

7.  up sec01一端

   ip netns exec 1-p20dp5pq1y ip link set sec01 up
```


### node节点操作
```bash
-   获取主机sec02 MAC地址与IP地址

1.  进入br0 ns 增加FDB表记录

bridge fdb append e6:db:09:43:42:b6 dev vxlan0 dst 192.168.60.120

2.  进入br0 ns 增加ARP表记录

ip neighbor add 10.0.1.252 lladdr e6:db:09:43:42:b6 dev vxlan0

```
