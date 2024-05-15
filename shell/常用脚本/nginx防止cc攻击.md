
```bash
#!/bin/bash
set -x
## 过滤日志文件中近5分钟日志，某个ip在5分钟内超过100次请求，会被认为异常访问.
## 异常访问会通过iptables 封禁源ip
## 如需解封，请手动执行iptables -D INPUT $rule_number 完成解封，删除规则，实时生效.

LOG_FILE=/var/log/nginx/access.log

## 拼接最近5分钟时间字符串，用于grep过滤
for i in {1..5}
do
        local_time=$(date -d "-${i} minutes" +'%Y:%H:%M')

        if [ -z "$a" ]; then
                a="${local_time}"
        else
                a="$a|${local_time}"
        fi
done
## 过滤出5分钟内的日志，并统计每个ip出现的次数，大于100进入ban_ips列表
ban_ips=$(grep -E "${a}" $LOG_FILE |awk '{print $1}' |sort |uniq -c| awk  '$1 > 100  {print $2}')

## 通过iptables封禁ban_ips


for ban_ip in $ban_ips
do
        if ! /usr/sbin/iptables -C INPUT -s $ban_ip -j DROP;then
                /usr/sbin/iptables -I INPUT -s $ban_ip -j DROP
                echo "$ban_ip 异常访问，已通过防火墙封禁."
        fi
done

```