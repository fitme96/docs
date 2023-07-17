curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
执行会安装
docker-ce docker-ce-cli containerd.io docker-compose-plugin docker-ce-rootless-extras docker-buildx-plugin

#### dockerd 设置代理

root@sec:~# mkdir -p /etc/systemd/system/docker.service.d/
root@sec:~# cat >> /etc/systemd/system/docker.service.d/http-proxy.conf << EOF
[Service]
Environment="HTTP_PROXY=http://192.168.60.49:7890/"
Environment="HTTPS_PROXY=http://192.168.60.49:7890/"
Environment="NO_PROXY=localhost,127.0.0.1,.example.com"
EOF
systemctl daemon-reload 
systemctl restart docker