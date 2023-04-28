脚本分析

bash_init.sh

##### 主机初始化
```
## 关闭selinux firewalld 并设置时区为上海
sed -i "s/SELINUX=enforcing$/SELINUX=disabled/g" /etc/selinux/config
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
systemctl stop firewalld.service
systemctl disable firewalld.service

## 生成本地源
mv /etc/yum.repos.d /et/yum.repos.d.bak
cp config/local.repo /etc/yum.repos.d/
echo "baseurl=$RPM_URL" >> /etc/yum.repos.d/local.repo


## ntfs-3g安装  d
yum install  -y wget gcc 
tar -zxvf package/ntfs-3g_ntfsprogs-2017.3.23.tgz -C package/
cd package/ntfs-3g_ntfsprogs-2017.3.23/
./configure && make && make install

### 设置hostname 增加hosts  
rm -r -f /etc/hostname
rm -r -f /etc/hosts
hostnamectl set-hostname controller
echo "$controller $controller_name" >> /etc/hosts
echo "$docker $docker_name" >> /etc/hosts
echo "127.0.0.1 localhost" >> /etc/hosts
localectl  set-locale LANG=en_US.UTF-8

## install redis
wget vim net-tools screen chrony etcd redis d
systemctl enable redis
systemctl start redis

##  d
yum install -y python-openstackclient mysql-server python2-PyMySQL openstack-selinux mysql-devel rabbitmq-server libxslt-devel memcached python-memcached openstack-keystone httpd mod_wsgi openstack-glance python-pip git python-devel libffi-devel gcc openssl-devel openstack-nova-api openstack-nova-conductor openstack-neutron openstack-neutron-ml2 openstack-dashboard python-setuptools openstack-nova-console openstack-nova-novncproxy  openstack-neutron-openvswitch ebtables openstack-nova-scheduler openstack-nova-placement-api openstack-nova-compute

systemctl restart httpd.service

## 编译python3.7.4 d 
yum install -y nginx supervisor zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel gdbm-devel db4-devel libpcap-devel xz-devel libffi-devel

tar -zxvf package/Python-3.7.4.tgz -C package/
cd package/Python-3.7.4/
./configure --prefix=/usr/local/python3
make && make install
ln -s /usr/local/python3/bin/python3.7 /usr/bin/python3.7
ln -s /usr/local/python3/bin/pip3.7 /usr/bin/pip3.7

### remove docker  d
yum -y  remove docker \
docker-client \
docker-client-latest \
docker-common \
docker-latest \
docker-latest-logrotate \
docker-logrotate \
docker-selinux \
docker-engine-selinux \
docker-engine

## 安装docker d 
        yum install -y -q docker-ce
        yum install -y -q docker-ce-rootless-extras
        sudo yum -y install conntrack-tools
        sudo yum install -y  targetcli python-keystone
        sudo yum -y install libvirt
        sed -i "s/\ExecStart=\/usr\/bin\/dockerd/ExecStart=\/usr\/bin\/dockerd -H tcp:\/\/127.0.0.1:2375 -H unix:\/\/var\/run\/docker.sock/g" /usr/lib/systemd/system/docker.service
        sudo systemctl daemon-reload
        sudo systemctl restart docker
        sudo systemctl enable docker

        yum install -y python-devel
        yum install -y python-pip
        yum -y install htop
        systemctl start libvirtd
        systemctl enable libvirtd
        systemctl start mysqld

## 



cpupower frequency-set -g performance

/usr/bin/mysqladmin -uroot -p${mypassw} password $MYSQL_ROOT
source /adpdata/trainsec/config/config_pass
systemctl restart mysqld



source /adpdata/trainsec/config/config_pass
mysql -u root -p$MYSQL_ROOT -e"
show variables like 'max_connections';
quit"
shutdown -r now

```



