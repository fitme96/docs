
https://redis.io/download/
```
tar xf redis-7.2.4.tar.gz
cd redis-7.2.4

make -j4 && make PREFIX=/usr/local/redis install

/usr/local/redis/bin/redis-server -v

```


问题记录
redis 安装报错 jemalloc/jemalloc.h: No such file or directory
make distclean  && make   清楚上次编译残留文件，重新编译，参考如下文章
https://www.cnblogs.com/operationhome/p/10342258.html