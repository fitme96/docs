### 升级

单机2.1.0 -> 2.3.1, 使用内嵌数据库

```bash
cd ${nacos_home}
./bin/shutdown.sh

备份data
mkdir /usr/local/nacos.new/ && cp -a ${nacos_home}/data /usr/local/nacos.new/

unzip nacos-server-2.3.1.zip -d /usr/local/nacos.new/

/usr/local/nacos.new/bin/startup.sh -m standalone



```