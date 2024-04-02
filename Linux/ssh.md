
升级openssh9.7p1 需要openssl >=1.1.1

注意
1. 可以开启telnet登录，这样ssh升级失败可以通过telnet处理
2. 不要移除旧版本openssl，很多软件依赖openssl
#### 安装telnet-server

```bash
yum install -y telnet-server

systemctl start telnet.socket


echo "pts/3" >> /etc/securetty

使用telnet工具测试登录成功后再操作下面步骤

```


#### 升级openssl
```
yum install gcc openssl-devel zlib-devel


curl https://www.openssl.org/source/old/1.1.1/openssl-1.1.1w.tar.gz -o openssl-1.1.1w.tar.gz

tar xf openssl-1.1.1w.tar.gz && cd openssl-1.1.1w

./config --prefix=/usr/local/ssl --openssldir=/usr/local/ssl shared

make -j4 && make install
sh -c 'echo "/usr/local/ssl/lib" >> /etc/ld.so.conf.d/openssl-1.1.1.conf'

ldconfig -v |grep ssl

```
#### 升级openssh
```shell
yum install pam-devel -y
wget https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.7p1.tar.gz

tar xf openssh-9.7p1.tar.gz && cd openssh-9.7p1

./configure --prefix=/usr/local/openssh.97p1  --with-ssl-dir=/usr/local/ssl --with-pam

make -j4 && make install
/usr/local/openssh.97p1/bin/ssh -V
mv /usr/sbin/sshd /usr/sbin/sshd.old
ln -sv /usr/local/openssh.97p1/sbin/sshd  /usr/sbin/sshd


sed -i.bak 's#notify#simple#g' /lib/systemd/system/sshd.service

使用如下systemd配置文件,主要修改Type为simple
[Unit]
Description=OpenSSH server daemon
Documentation=man:sshd(8) man:sshd_config(5)
After=network.target sshd-keygen.service
Wants=sshd-keygen.service

[Service]
Type=simple
EnvironmentFile=/etc/sysconfig/sshd
ExecStart=/usr/sbin/sshd -D $OPTIONS
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=on-failure
RestartSec=42s

[Install]
WantedBy=multi-user.target


systemctl daemon-reload

修改sshd_config UsePAM 改为yes ，因为编译时开启了pam模块，必须要打开这个参数

注意： 配置文件在这里 /usr/local/openssh.97p1/etc/

mv /usr/local/openssh.97p1/etc/sshd_config /usr/local/openssh.97p1/etc/sshd_config.bak
cp /etc/ssh/sshd_config /usr/local/openssh.97p1/etc/

sed -i.bak 's/GSSAPIAuth.*/#&/g' /usr/local/openssh.97p1/etc/sshd_config
sed -i.bak 's/GSSAPICleanup.*/#&/g' /usr/local/openssh.97p1/etc/sshd_config
systemctl restart sshd
```


#### 移除telnet服务
```
yum -y remove telnet-server
```