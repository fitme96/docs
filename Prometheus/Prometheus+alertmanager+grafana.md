docker-compose.yml
```yaml
version: '2.1'

services:
  prom:
    image: prom/prometheus:latest
    container_name: prom
    volumes:
      - prom_conf:/etc/prometheus
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
      - '--web.enable-lifecycle'
    ports:
      - 9090:9090
  grafana:
    image: grafana/grafana
    container_name: grafana
    restart: always
    volumes:
      - grafana:/var/lib/grafana
    ports:
      - 3000:3000
  alertmanager:
    image: bitnami/alertmanager:latest
    restart: always
    volumes:
      - alertmanager:/opt/bitnami/alertmanager/conf/
    ports:
      - 9093:9093
volumes:
  prom_conf:
  grafana:
  alertmanager:
  prometheus_data:

```

prometheus配置文件prometheus.yml
```
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

    static_configs:
      - targets: ["localhost:9090"]

```

alertmanager 配置文件config.yml

```yaml
cat config.yml 
global:
  resolve_timeout: 1m
route:
  group_by: ['alertname']
  group_wait: 5s
  group_interval: 5m
  repeat_interval: 3h
  receiver: 'web.hook'
receivers:
  - name: 'web.hook'
    webhook_configs:
      - url: 'http://192.168.230.12:8001/alert'
        send_resolved: false
```

