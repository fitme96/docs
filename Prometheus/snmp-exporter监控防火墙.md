部署snmp-exporter
- 通过snmp-exporter暴露给prometheus，需要snmp.yml

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

prometheus.yml 增加如下
```yaml
  - job_name: "snmp-firewall"
    static_configs:
      - targets:
        - "192.168.255.1"
    metrics_path: "/snmp"
    params:
      module: [firewall]
    scrape_interval: 10s
    scrape_timeout: 5s
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 192.168.230.20:9116  # The SNMP exporter's real hostname:port.

```

snmp告警规则
```
cat rules/snmp.rules 
groups:
- name: firewall
  rules:
  - alert: system11
    expr: system11{system11!="", system11=~"^(8[9]|[9][0-9]|100)$"}
    for: 10s
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} 设备温度过高"
      description: "{{ $labels.instance }} 温度{{ $labels.system11}}"
  - alert: cpuUsage
    expr: cpuUsage{cpuUsage!="", cpuUsage=~"^(8[9]|[9][0-9]|100)$"}
    for: 10s
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} CPU使用率过高"
      description: "{{ $labels.instance }} CPU使用率为{{ $labels.cpuUsage}}"
  - alert: memUsage
    expr: memUsage{memUsage!="", memUsage=~"^(8[9]|[9][0-9]|100)$"}
    for: 10s
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} 内存使用率过高"
      description: "{{ $labels.instance }} 内存使用率为{{ $labels.memUsage}}"

```

### 生成snmp.yml

1. 本地构建生成
依赖go环境
```shell
git clone https://github.com/prometheus/snmp_exporter.git github.com/snmp-exporter

cd github.com/snmp-exporter/generator


go env -w GOPROXY=https://goproxy.cn,direct

go get github.com/prometheus/snmp_exporter/generator

go build

cp generator.yml{,.bak}

```
2. 使用docker镜像生成
```shell
docker run --entrypoint bash  -ti  prom/snmp-generator

rm -f  /etc/apt/sources.list.d/*

echo "deb http://mirrors.ustc.edu.cn/debian stable main contrib non-free non-free-firmware
# deb-src http://mirrors.ustc.edu.cn/debian stable main contrib non-free non-free-firmware
deb http://mirrors.ustc.edu.cn/debian stable-updates main contrib non-free non-free-firmware
# deb-src http://mirrors.ustc.edu.cn/debian stable-updates main contrib non-free non-free-firmware

# deb http://mirrors.ustc.edu.cn/debian stable-proposed-updates main contrib non-free non-free-firmware
# deb-src http://mirrors.ustc.edu.cn/debian stable-proposed-updates main contrib non-free non-free-firmware" > /etc/apt/sources.list.d/nofree.list

apt update
apt install -y snmp-mibs-downloader


```

```shell
root@e90cdb0064d4:~# cat /etc/snmp/snmp.conf  
# As the snmp packages come without MIB files due to license reasons, loading
# of MIBs is disabled by default. If you added the MIBs you can reenable
# loading them by commenting out the following line.
mibs :
mibs +SNMPv2-MIB # +ALL 加载所有 +SNMPv2-MIB是模块名，通过下文a指令查询

# If you want to globally change where snmp libraries, commands and daemons
# look for MIBS, change the line below. Note you can set this for individual
# tools with the -M option or MIBDIRS environment variable.
#
mibdirs /usr/share/snmp/mibs:/usr/share/snmp/mibs/iana:/usr/share/snmp/mibs/ietf:/usr/share/snmp/mibs-downloader/mibiana:/usr/share/snmp/mibs-downloader/mibrfcs # 增加mib目录

```
a.
```bash
root@e90cdb0064d4:~# snmptranslate  -IR sysName
SNMPv2-MIB::sysName
```

b. generator.yml
```yaml
---
auths:
  public_v1:
    version: 1
  public_v2:
    version: 2

modules:
  switch:
    walk:
      - sysUpTime
      - sysName
```
c. snmp.yml
```yml
root@e90cdb0064d4:~# cat snmp.yml 
# WARNING: This file was auto-generated using snmp_exporter generator, manual changes will be lost.
auths:
  public_v1:
    community: public
    security_level: noAuthNoPriv
    auth_protocol: MD5
    priv_protocol: DES
    version: 1
  public_v2:
    community: public
    security_level: noAuthNoPriv
    auth_protocol: MD5
    priv_protocol: DES
    version: 2
modules:
  switch:
    get:
    - 1.3.6.1.2.1.1.3.0
    - 1.3.6.1.2.1.1.5.0
    metrics:
    - name: sysUpTime
      oid: 1.3.6.1.2.1.1.3
      type: gauge
      help: The time (in hundredths of a second) since the network management portion
        of the system was last re-initialized. - 1.3.6.1.2.1.1.3
    - name: sysName
      oid: 1.3.6.1.2.1.1.5
      type: DisplayString
      help: An administratively-assigned name for this managed node - 1.3.6.1.2.1.1.5

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
