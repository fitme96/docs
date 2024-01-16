```bash
yum install -y vsftpd ftp


[root@20231017-zhongjianbin-openeuleros ~]# grep -v '^#' /etc/vsftpd/vsftpd.conf 
# 允许匿名用户
anonymous_enable=NO
# 系统本地用户
local_enable=YES
# 
write_enable=YES
local_umask=022
# 本地根目录
local_root=/root/vmutils/target/ 
# 允许chroot
allow_writeable_chroot=YES

dirmessage_enable=YES
xferlog_enable=YES
connect_from_port_20=YES
xferlog_std_format=YES
chroot_local_user=YES
listen=NO
listen_ipv6=YES

pam_service_name=vsftpd
userlist_enable=YES


ftp localhost

chroot目录在/root下时需要chmod +x /root  ，root目录默认不允许其他用户执行x

```


- 匿名用户使用的登陆名为ftp或anonymous，口令为空；匿名用户不能离开匿名用户家目录/var/ftp,且只能下载不能上传
- 本地用户的登录名为本地用户名，口令为此本地用户的口令；本地用户可以在自 己家目录中进行读写操作；本地用户可以离开自家目录切换至有权限访问的其他目录，并在权限允许的情况下进行上传/下载