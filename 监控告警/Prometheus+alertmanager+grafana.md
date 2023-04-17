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
