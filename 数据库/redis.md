
https://redis.io/download/
```
tar xf redis-7.2.4.tar.gz
cd redis-7.2.4

make -j4 && make PREFIX=/usr/local/redis install

/usr/local/redis/bin/redis-server -v

```
