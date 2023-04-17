storcli安装包[storcli download](https://docs.broadcom.com/docs/007.0709.0000.0000_Unified_StorCLI.zip),下载解压后，安装rpm包
```shell
[root@localhost ~]# cd Unified_storcli_all_os
[root@localhost Unified_storcli_all_os]# ls
EFI  FreeBSD  JSON-Schema  Linux  Linux-PPC  Ubuntu  VMwareOP  Windows
[root@localhost Unified_storcli_all_os]# ls Linux
license.txt  LINUX_Readme.txt  splitpackage.sh  storcli-007.0709.0000.0000-1.noarch.rpm
```
检查RAID磁盘有无坏道,所有磁盘查看Media Other 统计数，不大于不用换磁盘
```shell
[root@localhost ~]# /opt/MegaRAID/storcli/storcli64 /c0 /eall /sall show all

Media Error Count = 0
Other Error Count = 0
```
检查RAID磁盘smart信息, 其中5 为DID通过storcli命令获取
```shell
/opt/MegaRAID/storcli/storcli64 /c0 show all 获取DID

[root@localhost ~]# smartctl -A -d megaraid,5 /dev/sda
smartctl 7.0 2018-12-30 r4883 [x86_64-linux-3.10.0-1160.11.1.el7.x86_64] (local build)
Copyright (C) 2002-18, Bruce Allen, Christian Franke, www.smartmontools.org

=== START OF READ SMART DATA SECTION ===
Current Drive Temperature:     31 C
Drive Trip Temperature:        60 C

Manufactured in week 01 of year 2019
Specified cycle count over device lifetime:  10000
Accumulated start-stop cycles:  83
Specified load-unload count over device lifetime:  300000
Accumulated load-unload cycles:  1405
Elements in grown defect list: 0   #有无坏道

```
### 参考
[storcli的安装](https://ahelpme.com/hardware/lsi/install-the-new-storcli-to-manage-lsi-avago-broadcom-megaraid-controller-under-centos-7/)