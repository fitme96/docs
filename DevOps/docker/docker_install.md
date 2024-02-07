## 安装

```bash
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
# 执行会安装
docker-ce docker-ce-cli containerd.io docker-compose-plugin docker-ce-rootless-extras docker-buildx-plugin


# 安装docker-compose
 curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# 插件方式安装，docker子命令
DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.24.5/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose

```


## dockerd 设置代理
```bash
root@sec:~# mkdir -p /etc/systemd/system/docker.service.d/
root@sec:~# cat >> /etc/systemd/system/docker.service.d/http-proxy.conf << EOF
[Service]
Environment="HTTP_PROXY=http://192.168.60.49:7890/"
Environment="HTTPS_PROXY=http://192.168.60.49:7890/"
Environment="NO_PROXY=localhost,127.0.0.1,.example.com"
EOF
systemctl daemon-reload 
systemctl restart docker
```
