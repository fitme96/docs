### Netfilter/Iptables 信息包过滤系统可以当成一个整体，netfilter是内核的模块实现，iptables是对上层操作工具

#### 表
- raw
- mangle
- filter
- nat

####  链
每个链都有多条规则，且有默认规则，规则从第一条开始匹配，--line-numbers参数可以查看规则号，匹配未中时执行默认规则
- INPUT
- OUTPUT
- FORWARD
- PREROUTING
- POSTROUTING
#### 动作

（注意：一旦执行完规则的动作，就结束整个规则链的匹配，不继续往下匹配其他规则了。）

1. Accept：接受报文  
2. Drop：丢弃报文 
3. SNAT：把报文的源地址改掉 （需要指定改为什么IP）
4. DNAT：把报文的目的地址改掉
5. Masquerade：高级版的SNAT，把报文的源地址改掉，改为发送网卡的IP地址（自动识别改成什么IP）
6. Redirect：高级版的DNAT，把报文的目的地址改为接收网卡的IP地址。同时可以修改目的端口（其实就是重定向给本机的某个端口）
7. 其他名字：那就代表跳转到另外一条规则链继续匹配。。。（这里的名字就等于指定的目标规则链的名字）

### 规则命中次数
```shell
iptables -Z 清空统计数
iptables -nvL

```

### iptables调试

```shell

iptables-save

## 开启log
modprobe nf_log_ipv4
sysctl net.netfilter.nf_log.2 

## 通过raw插入trace
iptables -t raw -A OUTPUT -p TCP -j TRACE
iptables -t raw -A PREROUTING -p TCP -j TRACE
iptables -nvL -t raw
## 通过syslog过滤分析iptables过滤
tail  -f /var/log/syslog |grep "192.168.65.8"
```

规则
```bash
默认规则设置
iptables -P FORWARD DROP
iptables -P INPUT ACCEPT

iptables -A  FORWARD -j ACCEPT 

```