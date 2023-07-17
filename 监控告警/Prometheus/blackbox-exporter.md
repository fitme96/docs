
探测端增加如下服务
```yaml

  blackbox-exporter:
    image: prom/blackbox-exporter
    command:
      - '--config.file=/config/blackbox.yml'
    volumes:
      - './blackbox.yml:/config/blackbox.yml'
```


prometheus.yml增加如下
```yaml

## blackbox
  - job_name: 'blackbox'
    metrics_path: /probe
    params:
      module: [http_2xx]  # Look for a HTTP 200 response.
    file_sd_configs:
      - files:
          - "httpstatustargets.yml"
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: blackbox-exporter:9115  # The blackbox exporter's real hostname:port.
```

httpstatustargets.yml

```yaml

- targets:
    - "http://192.168.60.197:3000"
  labels:
    env: sec
```


映射blackbox.yml
```

root@Tools:/data/prometheus# cat blackbox.yml 
modules:
  http_2xx:
    prober: http
    http:
      method: GET
  http_post_2xx:
    prober: http
    http:
      method: POST
```