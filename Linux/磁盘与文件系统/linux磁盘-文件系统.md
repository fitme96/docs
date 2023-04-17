
安装smartctl
```shell
yum install -y smartmontools
```
查看硬盘是否打开smart支持
```shell
smartctl -i /dev/sda
```

如下命令查看硬盘的健康状况：  
```shell
smartctl -H /dev/sdb  
  
=== START OF READ SMART DATA SECTION ===  
SMART overall-health self-assessment test result: PASSED  
  
请注意result后边的结果：PASSED，这表示硬盘健康状态良好，如果这里显示Failure，那么最好立刻给服务器更换硬盘。  
  
SMART只能报告磁盘已经不再健康，但是报警后还能继续运行多久是不确定的。通常，SMART报警参数是有预留的，磁盘报警后，不会当场坏掉，一般能坚持一段时间  
有的硬盘SMART报警后还继续跑了好此年，有的硬盘SMART报错后几天就坏了。
```



磁盘坏道坏块检测
sudo badblocks -s -v /dev/sdb3

文件系统修复

xfs

ext2/3/4