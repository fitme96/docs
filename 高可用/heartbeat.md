---
title: heartbeat
type: 高可用
 ---



## hearbeat 安装
```bash
apt install -y heartbeat
cp /usr/share/doc/heartbeat/{authkeys,ha.cf.gz,haresources.gz} /etc/ha.d
gzip -d ha.cf.gz haresources.gz

### vim authkeys
增加
auth 3
3 md5 Hello!
END

chmod 600 authkeys

### vim haresources
增加
sec 192.168.180.224/24/ens18:1 yulei
END
说明: sec 主节点主机名, 192.168.180.224是VIP，ens18:1 是VIP设备名 ， yulei 是监控服务（yulei需要放入/etc/init.d或者/etc/ha.d/resources.d/下，并且实现stop 和 start函数

### vim ha.cf
最后追加
debugfile /var/log/ha-debug
logfile /var/log/ha-log
logfacility local0
keepalive 2
deadtime 30
warntime 10
initdead 60
udpport 694
ucast ens19 192.168.223.23
auto_failback off
node sec
node sec2
ping 192.168.180.254
respawn root /usr/lib/heartbeat/ipfail
apiauth ipfail gid=root  uid=root
END


### 同步上面三个文件到从节点
修改authkeys权限
chmod 600 authkeys

修改ha.cf
ucast ens19 192.168.223.23 为 192.168.223.22 

systemctl start heartbeat



```


