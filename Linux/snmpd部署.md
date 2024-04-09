#### 在设备启动snmp代理

snmp 客户端软件  snmpd 代理  libsnmp-dev 工具软件  snmp-mibs-downloader 下载本地库mib库软件
apt install snmp snmpd libsnmp-dev snmp-mibs-downloader
默认监听回环地址，更改为所有
agentaddress  127.0.0.1,[::1]  ->  agentaddress  0.0.0.0,[::1]
获取更多节点信息
vim /etc/snmp/snmpd.conf
-view   systemonly  included   .1.3.6.1.2.1.1    
-view   systemonly  included   .1.3.6.1.2.1.25.1 
+view   systemonly  included   .1                 

systemctl restart snmpd
##### 更改community

community默认是public

vim /etc/snmp/snmpd.conf
rocommunity  public default -V systemonly
rocommunity6 public default -V systemonly
public改为cc
rocommunity  cc default -V systemonly
rocommunity6 cc default -V systemonly

#### 客户端获取数据
```
snmpget -v 2c -c cc 192.168.190.108 SNMPv2-MIB::sysName.0                               
SNMPv2-MIB::sysName.0 = STRING: node0
```


#### V3版本

v3有三种方式/ 不认证不加密、认证不加密、认证加密

创建用户时关闭snmpd服务
systemctl stop snmpd
fitme123认证密码 cunkai123 私有密码 cunkai 用户
net-snmp-config --create-snmpv3-user -ro -A fitme123 -X cunkai123 -a SHA -x AES cunkai


v3版本获取数据

#### 参考

[SNMPv3]([SNMPv3 配置及snmpwalk命令信息获取 - quliuliu2013 - 博客园](https://www.cnblogs.com/quliuliu2013/p/10056933.html))


### MIB下载
[华为常用MIB节点](https://support.huawei.com/enterprise/zh/doc/EDOC1100035605/ee5270d1)
