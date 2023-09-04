- nsenter可进入命空间，常见使用于使用主机命令调试容器
```shell
nsenter [options] [program [arguments]]

options:
-t, --target pid：指定被进入命名空间的目标进程的pid
-m, --mount[=file]：进入mount命令空间。如果指定了file，则进入file的命令空间
-u, --uts[=file]：进入uts命令空间。如果指定了file，则进入file的命令空间
-i, --ipc[=file]：进入ipc命令空间。如果指定了file，则进入file的命令空间
-n, --net[=file]：进入net命令空间。如果指定了file，则进入file的命令空间
-p, --pid[=file]：进入pid命令空间。如果指定了file，则进入file的命令空间
-U, --user[=file]：进入user命令空间。如果指定了file，则进入file的命令空间
-G, --setgid gid：设置运行程序的gid
-S, --setuid uid：设置运行程序的uid
-r, --root[=directory]：设置根目录
-w, --wd[=directory]：设置工作目录

如果没有给出program，则默认执行$SHELL。
```

## 使用
```shell
root@ck-test-65:~# docker ps
CONTAINER ID   IMAGE                       COMMAND                  CREATED       STATUS       PORTS                                       NAMES
8a4a167bc951   nginx                       "/docker-entrypoint.…"   6 days ago    Up 6 days    80/tcp                                      serene_bohr
94a3f541e3f4   prom/node-exporter:latest   "/bin/node_exporter …"   3 weeks ago   Up 2 weeks   0.0.0.0:9100->9100/tcp, :::9100->9100/tcp   node_exporter
root@ck-test-65:~# docker exec -ti 8a4a bash
root@8a4a167bc951:/# ping baidu.com
bash: ping: command not found
root@8a4a167bc951:/# 
exit
root@ck-test-65:~# docker inspect 8a4 -f '{{.State.Pid}}'
33727
root@ck-test-65:~# nsenter -n -t 33727
root@ck-test-65:~# ping baidu.com
ping: baidu.com: Temporary failure in name resolution
root@ck-test-65:~# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
11: eth0@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

```