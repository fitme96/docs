
升级openssh9.7p1

需要openssl >=1.1.1

```
yum install gcc openssl-devel zlib-devel


curl https://www.openssl.org/source/old/1.1.1/openssl-1.1.1w.tar.gz -o openssl-1.1.1w.tar.gz

tar xf openssl-1.1.1w.tar.gz && cd openssl-1.1.1w

./config --prefix=/usr/local/ssl --openssldir=/usr/local/ssl shared



```