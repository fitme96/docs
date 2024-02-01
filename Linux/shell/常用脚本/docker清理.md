```shell
#!/bin/bash

## 判断机器是否存在docker

docker version > /dev/null 2>&1   && echo "即将清理docker" || echo "此机器未安装docker"

## 清理退出的docker 容器 
## 清理未命名的镜像，释放磁盘空间
con=$(docker ps -a | grep "Exited" -c)
if [ "${con}" -ne 0 ];then
       docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker stop
       docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker rm
fi
docker images | grep none | grep -v REPOSITORY | awk '{print $3 }' | xargs docker rmi > /dev/null 2>&1

echo "docker none镜像，Exited容器清理完成"

```