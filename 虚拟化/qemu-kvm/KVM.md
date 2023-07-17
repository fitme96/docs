## 测试环境
```shell
root@kvm:~/certs# cat /etc/issue
Ubuntu 20.04.3 LTS \n \l


root@kvm:~# libvirtd -V
libvirtd (libvirt) 6.0.0
```

### 安装libvirtd

```shell
# 更改apt源

sed -i 's@//.*archive.ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list

systemctl stop ufw
systemctl disable ufw
# 判断KVM虚拟化扩展
apt install -y cpu-checker

root@kvm:~/certs# kvm-ok 
INFO: /dev/kvm exists
KVM acceleration can be used


# 安装kvm libvirtd
apt update
apt install qemu-kvm libvirt-daemon-system

adduser $USER libvirt
```



### libvirtd TLS配置

```shell
# 安装certtool工具管理证书
apt-get install gnutls-bin
mkdir ~/certs && cd ~/certs/
# 生成ca私钥
certtool --generate-privkey > cakey.pem


# 生成ca证书
cat >> ca.info <<EOF
cn = Sec
ca
cert_signing_key
EOF

certtool --generate-self-signed --load-privkey cakey.pem --template ca.info --outfile cacert.pem
# 检查ca证书
certtool -i --infile cacert.pem

# 生成服务端私钥

certtool --generate-privkey > serverkey.pem

cat >> server.info <<EOF 
organization = Sec
cn = kvm
tls_www_server
encryption_key
signing_key
EOF 
# 生成服务端证书
certtool --generate-certificate --load-privkey serverkey.pem  --load-ca-certificate cacert.pem --load-ca-privkey cakey.pem --template  server.info --outfile servercert.pem

# 生成客户端私钥

certtool --generate-privkey > clientkey.pem
cat >> client.info <<EOF
country = CN
state = Shaanxi
locality = Xian
organization = Sec
cn = client1
tls_www_client
encryption_key
signing_key
EOF
certtool --generate-certificate --load-privkey clientkey.pem  --load-ca-certificate cacert.pem --load-ca-privkey cakey.pem --template  client.info --outfile clientcert.pem

## 服务端放置ca证书及服务端证书 (/etc/libvirt/libvirtd.conf 中默认是如下位置，如果需要自定义更改libvirtd.conf "key_file" "ca_file" "cert_file" 三个参数即可 )
mkdir /etc/pki/CA/ 
cp cacert.pem /etc/pki/CA/
mkdir /etc/pki/libvirt/private -p 
cp servercert.pem /etc/pki/libvirt/
cp serverkey.pem /etc/pki/libvirt/private/

## 客户端放置ca证书及客户端证书
mkdir /etc/pki/CA/ 
cp cacert.pem /etc/pki/CA/
mkdir /etc/pki/libvirt/private -p 
cp clientcert.pem /etc/pki/libvirt/
cp clientkey.pem /etc/pki/libvirt/private/

systemctl stop libvirtd
systemctl start libvirtd-tls.socket
systemctl restart libvirtd


```

### 客户端远程调用测试

```shell
# 安装客户端工具
apt install libvirt-clients

# 在客户端增加hosts
192.168.65.49 kvm
# 在客户端使用hostname访问
root@client1:~# virsh -c qemu+tls://kvm/system list
 Id   Name     State
------------------------
 8    ubuntu   running

## 创建快照
root@client1:~# virsh -c qemu+tls://kvm/system snapshot-create-as ubuntu 525ubuntu-dd
Domain snapshot 525ubuntu-dd created
## 快照列表
root@client1:~# virsh -c qemu+tls://kvm/system snapshot-list ubuntu
 Name           Creation Time               State
-----------------------------------------------------
 1684997347     2023-05-25 06:49:07 +0000   running
 525-ubuntu     2023-05-25 05:40:57 +0000   running
 525ubuntu-dd   2023-05-25 06:50:58 +0000   running
## 恢复快照
root@client1:~# virsh -c qemu+tls://kvm/system snapshot-revert ubuntu 525-ubuntu


```

### libvirtd 配置认证
```shell

vim /etc/libvirt/libvirtd.conf
## 打开auth_tls,更改none为sasl
auth_tls = "sasl"

vim /etc/sasl2/libvirt.conf
## 打开mech_list ，更改为DIGEST-MD5
mech_list: DIGEST-MD5

# If using scram-sha-1 for username/passwds, then this is the file
# containing the passwds. Use 'saslpasswd2 -a libvirt [username]'
# to add entries, and 'sasldblistusers2 -f [sasldb_path]' to browse it
## 安装saslpasswd2工具
apt install sasl2-bin libsasl2-modules-gssapi-mit libsasl2-modules
## 增加user
saslpasswd2 -a libvirt ck
## 用户列表
sasldblistusers2 -f /etc/libvirt/passwd.db

systemctl stop libvirtd
systemctl start libvirtd-tls.socket

```


### 虚拟机创建

```shell
## 查看os列表
apt install -y libosinfo-bin
osinfo-query os
## 安装bridge及virt-install等工具
apt install -y  bridge-utils virtinst
## 配置网桥br0
root@kvm:~/certs# cat /etc/netplan/00-installer-config.yaml 
# This is the network config written by 'subiquity'
network:
  ethernets:
    eno1:
      dhcp4: false
    eno2:
      dhcp4: false
  # add configuration for bridge interface
  bridges:
    br0:
      interfaces: [eno1]
      dhcp4: false
      addresses: [192.168.65.49/24]
      gateway4: 192.168.65.254
      nameservers:
        addresses: [192.168.230.12]
      parameters:
        stp: false
      dhcp6: false
## netplan apply

## 创建磁盘
qemu-img create -f qcow2 /var/lib/libvirt/ubuntu.qcow2 100G
## 创建ubuntu KVM 虚拟机
virt-install     --virt-type kvm     --name ubuntu     --ram 4096     --cdrom=/var/lib/libvirt/ubuntu-20.04.3-live-server-amd64.iso    --disk path=/var/lib/libvirt/ubuntu.qcow2     --network bridge=br0     --graphics vnc,listen=0.0.0.0     --noautoconsole     --os-type=linux     --os-variant=ubuntu20.04
## 显示vnc 连接地址
root@kvm:~/certs# virsh vncdisplay ubuntu
:0
## 通过vnc Viewer 连接宿主机ip_addr:0


```
### 参考
- [认证](https://libvirt.org/auth.html)  
- [连接认证](https://libvirt.org/auth.html#ACL_server_username)  
- [远程连接](https://libvirt.org/remote.html#libvirtd-configuration-file)  
- [libvirtd manpages](https://manpages.ubuntu.com/manpages/focal/man8/libvirtd.8.html)
- https://bugzilla.redhat.com/show_bug.cgi?id=1589021
- https://ubuntu.com/server/docs/virtualization-libvirt
