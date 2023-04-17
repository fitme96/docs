部署node-exporter
```yaml
node-exporter:
    image: prom/node-exporter:latest
    container_name: node_exporter
    ports:
      - 9100:9100
    restart: always
    volumes:
      - /proc:/host/proc:ro
      - /sys:/host/sys:ro
      - /:/rootfs:ro
    command:
      - '--path.procfs=/host/proc'
      - '--path.rootfs=/rootfs'
      - '--path.sysfs=/host/sys'
      - '--collector.filesystem.mount-points-exclude=^/(sys|proc|dev|host|etc)($$|/)'
```

访问metricss
![[Pasted image 20230412151017.png]]

```yaml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
            - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  - rules/*.rules
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

### 增加如下通过watch文件动态增加监控主机
    static_configs:
      - targets: ["localhost:9090"]
  - job_name: "file-sec-sre"
    file_sd_configs:
      - files:
          - "targets.yml"
### END
```

```yaml
- targets:
    - "192.168.65.8:9100"
    - "192.168.230.20:9100"
    - "192.168.230.12:9100"
  labels:
    env: sec
```

增加告警策略
```yaml
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu_seconds_total{mode!='idle'}[5m]))) by (instance) > 0.85
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} CPU 使用率过高"
      description: "{{ $labels.instance }} CPU 使用率超过85% (current value: {{ $value }})"
  - alert: hostMemUsageAlert
    expr: (node_memory_MemTotal - node_memory_MemAvailable)/node_memory_MemTotal > 0.80
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} 内存使用率过高"
      description: "{{ $labels.instance }} 内存使用率超过80% (current value: {{ $value }})"
  - alert: hostDiskUsageAlert
    expr: (1 - node_filesystem_avail_bytes{fstype=~"ext4|xfs"} / node_filesystem_size_bytes{fstype=~"ext4|xfs"}) * 100 > 80
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} 磁盘使用率过高"
      description: "{{ $labels.instance }} 磁盘使用率超过80% (current value: {{ $value }})"
  - alert: hostUPAlert
    expr: up{job="file-sec-sre"}==0
    for: 10s
    labels:
      severity: page
    annotations:
      summary: "实例 {{ $labels.instance }} 不可用"
      description: "{{ $labels.instance }} 不可用"
```
