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

alertmanager 配置文件config.yml

```yaml
root@Tools:/var/lib/docker/volumes/prometheus_alertmanager/_data# cat config.yml 
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
        send_resolved: true
```

