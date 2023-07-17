deb包根据目录结构将二进制文件及配置文件解压放置

## dpkg名

打包目录结构示例

├── DEBIAN

│   ├── control

│   ├── postinst

│   └── postrm

├── lib

│   └── systemd

│   └── system

│   └── sec-portainer.service

└── usr

└── local

└── portainer

└── bin

1.  control 文件

Package: sec-portainer
Version: 1.0.0
Section: admin
Priority: optional
Architecture: amd64
Maintainer: cunkai cunkai@seclover.com
Description: yulei portainer

2.  postinst 文件(二进制文件放置后执行的命令

#!/bin/sh

set -e

if [ ! -d /data/portainer ];then
	mkdir /data/portainer -p >/dev/null 2>&1
	echo "create /data/portainer"
fi

3.  postrm 文件(卸载后动作

#!/bin/sh

set -e

4.  sec-portainer.service( systemd管理文件

[Unit]
Description=seclover portainer
Wants=network-online.target
After=network.target

[Service]
Type=exec
ExecStart=/usr/local/portainer/bin/portainer -d /data/portainer
Restart=always

[Install]
WantedBy=multi-user.target

5.  usr目录放置二进制文件