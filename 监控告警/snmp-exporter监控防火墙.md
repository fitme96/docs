部署snmp-exporter
- 通过snmp-exporter暴露给prometheus，需要snmp.yml，snmp.yml生成参考

```yaml
    snmp-exporter:
    image: prom/snmp-exporter
    command:
      - '--config.file=/etc/snmp_exporter/snmp.yml'
    volumes:
      - './snmp.yml:/etc/snmp_exporter/snmp.yml:ro'
    ports:
      - '9116:9116'
```

### 生成snmp.yml

依赖go环境
```shell
git clone https://github.com/prometheus/snmp_exporter.git github.com/snmp-exporter

cd github.com/snmp-exporter/generator

go env -w GOPROXY=https://goproxy.cn,direct

go get github.com/prometheus/snmp_exporter/generator

go build

cp generator.yml{,.bak}

```

修改 generator.yml
```yaml
modules:
  firewall:
    walk:
      - cpuUsage
      - memUsage
      - 1.3.6.1.4.1.15227.1.3.1.1
    auth:
      community: seclover
```

# 从厂商获取mib私库复制到mibs目录

```shell
mv /root/FW-VENUS-MIB\(1\).dat mibs/FW-VENUS-MIB.txt

export MIBDIRS=mibs

apt install -y unzip

make mibs

./generator generate

生成snmp.yml 用于snmp-exporter容器

```
