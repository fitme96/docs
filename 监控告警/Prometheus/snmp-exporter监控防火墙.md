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
