```shell
#!/bin/bash

## 判断机器是否存在docker
if ! docker version >/dev/null 2>&1 ;then
    echo "未安装docker,没有什么可做的"
    exit 0
else 
    echo "已安装dockerd,清理开始"
fi


## 清理退出的docker 容器 
## 清理未命名的镜像，释放磁盘空间
con=$(docker ps -a | grep "Exited" -c)
if [ "${con}" -ne 0 ];then
       docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker rm
fi
docker images | grep none | grep -v REPOSITORY | awk '{print $3 }' | xargs docker rmi > /dev/null 2>&1

echo "docker none镜像，Exited容器清理完成"

```