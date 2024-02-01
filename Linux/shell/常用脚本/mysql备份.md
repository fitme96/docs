```bash

#!/bin/bash
## 脚本使用需要传入mysql 用户名$1  密码$2
## $3用于指定mysql socket文件，用于非默认mysql安装
## 默认保留7天的备份数据
uname=$1
passwd=$2

# 备份目录及文件
bakDir="/data/database_backup"
nowDate=$(date +%Y-%m-%d-%H-%M-%S)
dumpFile="${bakDir}/all_db_${nowDate}.sql"

### 创建备份目录
if [ ! -d "${bakDir}" ]; then
  mkdir ${bakDir} -p
fi

## 传递参数3用于指定socket
mysqldump $3 -u${uname} -p${passwd} -A > ${dumpFile}

if [ $? -ne 0 ];then
   exit 1
fi

#定时清理
bakfile_number=7
find ${bakDir} -name "*.sql*" -mtime +${bakfile_number} | xargs -I {} rm -f {}

```