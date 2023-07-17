apt-get install dnsutils

```shell
nslookup baidu.com

dig baidu.com
## 指定DNS服务器解析
dig @223.5.5.5 baidu.com

```



这个 /etc/resolv.conf文件中，还可以配置一些高级参数。

- search：查询DNS域名时，会往你查询的域名尾部，额外补全的内容。
- ndots：控制补全的最大长度。