


## cloud-images
- https://cloud-images.ubuntu.com/
- https://cloud.centos.org/centos/ 
- https://cloud.debian.org/images/cloud/
- https://alt.fedoraproject.org/cloud/

## 创建虚拟机
```bash
## 创建虚拟机
qm create 101 --name centos7 --memory 4096 --net0 virtio,bridge=vmbr0
## 导入磁盘文件
qm importdisk 101 CentOS-7-x86_64-GenericCloud-2009.qcow2 local-lvm
## 设置磁盘总线类型为virtio
qm set 101 --virtio0 local-lvm:vm-101-disk-0
## 设置virtio0磁盘为第一引导设备
qm set 101 --boot c --bootdisk virtio0
## 添加cloudinit Driver设备
qm set 101 --ide2 local-lvm:cloudinit
## 虚拟机转换为模板
qm template 101

```

## 为用户创建虚拟机
### 通过cli为用户创建虚拟机
```bash
## 克隆101模板生成201虚拟机
qm clone 101 201 --name ck-test-65-211

```





