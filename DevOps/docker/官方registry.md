


# 增加TLS

## hostA 创建 registry    
```bash
mkdir -p certs
## 自签证书
openssl req \\n  -newkey rsa:4096 -nodes -sha256 -keyout certs/domain.key \\n  -addext "subjectAltName = DNS:hubseclover.com" \\n  -x509 -days 3650 -out certs/domain.crt
## 运行registry
sudo docker run -d \\n  --restart=always \\n  --name registry \\n  -v "$(pwd)"/certs:/certs \\n  -e REGISTRY_HTTP_ADDR=0.0.0.0:443 \\n  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \\n  -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \\n  -p 10443:443 \\n  registry:2
 
sudo docker tag registry:2 hubseclover.com/registry:2
sudo docker push hubseclover.com/registry:2  

scp certs/domain.crt root@$hostB:/root
 
```
## hostB 拉取镜像

```bash
echo 'hostA_IP hubseclover.com' >> /etc/hosts
mkdir /etc/docker/certs.d/hubseclover.com:10443/ -p
mv /root/domain.crt /etc/docker/certs.d/hubseclover.com:10443/ca.crt

docker pull hubseclover.com:10443/registry:2

```

## iso TODO
- 导入registry2镜像
- htpasswd
- 证书






