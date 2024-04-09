
```yaml
version: '2.1'

services:
  nacos:
    image: nacos/nacos-server:v2.3.1
    container_name: nacos
    environment:
      PREFER_HOST_MODE: hostname
      MODE: standalone
      NACOS_AUTH_ENABLE: true
      NACOS_AUTH_TOKEN: ZHNmc2RqZmxza2RqZmZzZGZzZGxmamRsc2ZzZGZzZGZqc2RramZsc2RqZmxzZGYKIA==
      NACOS_AUTH_IDENTITY_KEY: nacos
      NACOS_AUTH_IDENTITY_VALUE: nacos
    ports:
      - 8848:8848
      - 9848:9848
      - 7848:7848
    volumes:
      - nacos_conf:/home/nacos/conf
      - nacos_data:/home/nacos/data

```


### 一些说明
1. 开启鉴权必须设置 token.secret.key，容器方式通过环境变量NACOS_AUTH_TOKEN设置，可以 使用Base64编码大于等于32位字符生成
2. NACOS_AUTH_IDENTITY_VALUE 设置后不生效，可以进入web界面更改

### 关于Nacos升级风险及处理办法

[参考](https://nacos.io/zh-cn/blog/announcement-token-secret-key.html)