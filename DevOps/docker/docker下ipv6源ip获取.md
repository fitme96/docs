
## ipv6源ip
- 通过docker-proxy会重写客户端ip,ipv4可以获取正确远程地址是因为绕过了docker-proxy,NAT规则在代理前，docker并没有处理ipv6 nat，导致开启ipv6后不能正确获取远程地址。


- [docker-ipv6nat](https://github.com/robbertkl/docker-ipv6nat)应该可以解决问题，但是需要引入外部工具
- docker内部目前合并进来了，ipv6nat还处于实现性,在daemon.json增加如下:（测试来看iptables规则应该有问题，导致外部不能正确访问容器，待排查。)
```json
  "ipv6": true,
  "fixed-cidr-v6": "2001:db8:abc1::/64",
+ "experimental": true,
+ "ip6tables": true

```
## 参考
- https://github.com/robbertkl/docker-ipv6nat/issues/65
- https://github.com/caddyserver/caddy/issues/4339
- https://www.moyufangge.com/2020-12/docker-real-ip/
- https://github.com/moby/moby/issues/17666


