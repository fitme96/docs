1.  TASK ERROR: activating LV ‘pve/data’ failed: Activation of logical volume pve/data is prohibited while logical volume pve/data_tmeta is active.
```shell
lvchange -an pve/data_tmeta
lvchange -an pve/data_tdata
lvchange -ay pve

```


[参考](https://blog.csdn.net/feitianyul/article/details/125417765)

2. 虚拟机挂了，如何备份系统磁盘数据
```shell
apt install -y kpartx

kpartx -a /dev/pve/vm-104-disk-0
mount /dev/mapper/pve-vm--104--disk--0p1 /mnt/
kpartx -d /dev/pve/vm-104-disk-0

## 更简单干净的做法
guestfish -i -a /dev/pve/vm-104-disk-0 /mnt

```
[参考](https://backdrift.org/mounting-a-file-system-on-a-partition-inside-of-an-lvm-volume)

