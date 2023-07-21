```shell
# 查看指定日期
journalctl --since "2022-01-01" --until "2022-01-02"
```

systemd-analyze命令用于分析系统启动时间。以下是一些常用的systemd-analyze命令：

- 显示系统启动时间：

```shell
systemd-analyze
```


- 显示各个服务单元的启动时间：

```shell
systemd-analyze blame
```