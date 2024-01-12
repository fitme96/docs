```bash
ntp官方下载地址：https://www.ntp.org/downloads/
wget https://archive.ntp.org/ntp4/ntp-4.2/ntp-4.2.8p17.tar.gz

tar xf ntp-4.2.8p17.tar.gz && cd ntp-4.2.8p17

yum install gcc gcc-c++ openssl-devel libstdc++* libcap*

./configure --prefix=/usr --bindir=/usr/sbin --enable-all-clocks --enable-parse-clocks --docdir=/usr/share/doc/ntp-4.2.8p17 --enable-all-clocks --enable-parse-clocks --enable-clockctl CFLAGS="-fPIC"

make && make install

cat  > /usr/lib/systemd/system/ntpd.service << EOF
[Unit]
Description=Network Time Service
After=syslog.target ntpdate.service sntp.service
[Service]
Type=forking
EnvironmentFile=-/etc/sysconfig/ntpd
ExecStart=/usr/sbin/ntpd -u root:root \$OPTIONS
PrivateTmp=true
[Install]
WantedBy=multi-user.target
EOF


vim /etc/ntp.conf
EOF
driftfile /var/lib/ntp/drift
restrict default nomodify notrap nopeer noepeer noquery
restrict source nomodify notrap noepeer noquery
restrict 127.0.0.1 
restrict ::1
tos maxclock 5
includefile /etc/ntp/crypto/pw
keys /etc/ntp/keys
server ntp.aliyun.com
EOF

systemctl daemon-reload
systemctl start ntpd
systemctl enable ntpd

ntpq -p
[root@dict-water2 ~]# ntpq -p

     remote           refid      st t when poll reach   delay   offset  jitter
==============================================================================
*203.107.6.88    100.107.25.114   2 u   16   64  377    5.821   +3.900   2.808

```