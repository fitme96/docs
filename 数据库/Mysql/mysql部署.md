```bash
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v mysql_data:/var/lib/mysql mysql:8.0.32

mysql8.0允许root远程登录

alter user 'root'@'%' identified with mysql_native_password by '123456';

flush privileges;

```




