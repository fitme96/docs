1.  异常断电导致dockerd无法启动

原因：容器network丢失

解决：

mv /var/lib/docker/network /tmp/
systemctl start docker
容器不能正常启动
compose 启动得容器直接通过 docker compose up --force-recreate CONTAINER_NAME 重建即可
docker run 启动的容器，需要通过inspect找到启动参数完成启动

2.  异常断电导致postgres启动失败

原因:Corruption in "pgsql/data/pg_logical/replorigin_checkpoint" file

解决:

Move out the "pgsql/data/pg_logical/replorigin_checkpoint" file to somewhere else for backup purpose.  
Restart the PostgreSQL database.

cannot start with following error postgresql* log files:
replication checkpoint has wrong magic number xxxxxx instead of yyyyyy