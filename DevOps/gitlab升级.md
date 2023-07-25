### 当前环境说明

-   gitlab安装方式使用dockerhub gitlab/gitlab:13.12.15-ce.0镜像
-   目标版本15.0.2

### 先决条件

-   提前拉取所有升级过程中的中间版本

### 升级步骤

1.  升级通知
2.  源gitlab生成备份数据并停止源gitlab
3.  新建一台同版本gitlab
4.  将备份数据恢复(采用gitlab-rake gitlab:backup:create备份数据需要手动备份gitlab.rb ，gitlab-secrets.json并一同恢复)
5.  根据升级路线做升级，检查

#### 升级路线

-   13.12.15->14.0.12->14.3.6->14.6.2->14.9.5->14.10.5->15.0.2
-   [官方推荐升级路线](https://docs.gitlab.com/ee/update/#1420)

### 详细升级过程:

1.  启动新的gitlab容器(变更数据映射目录）

version: '3.6'
services:
gitlabnew:
 image: 'gitlab/gitlab-ce:13.12.15-ce.0'
 hostname: 'git.seclover.com'
 restart: always
 environment:
   GITLAB_OMNIBUS_CONFIG: |
     external_url 'https://git.seclover.com'
     # Add any other gitlab.rb configuration here, each on its own line
 ports:
   - '80:80'
   - '443:443'
   - '22:22'
 volumes:
   - '/data/gitlab/config:/etc/gitlab'
   - '/data/gitlab/logs:/var/log/gitlab'
   - '/data/gitlab/data:/var/opt/gitlab'

2. 源gitlab容器获取备份包
```bash
docker exec -ti gitlab bash
gitlab-rake gitlab:backup:create
docker cp gitlab:/var/opt/gitlab/backups/1658977715_2022_07_28_13.12.15_gitlab_backup.tar .

3.  恢复备份包到新gitlab容器

docker cp 1658977715_2022_07_28_13.12.15_gitlab_backup.tar gitlabnew:/var/opt/gitlab/backups/
docker exec -ti gitlabnew bash
chmod 777 /var/opt/gitlab/backups/1658977715_2022_07_28_13.12.15_gitlab_backup.tar
gitlab-rake gitlab:backup:restore BACKUP=1658977715_2022_07_28_13.12.15

4.  升级13.12.15 -> 14.0.12  
    修改compose文件直接up
5.  14.0.12  -> 14.3.6

gitlab-rake db:migrate
gitlab-psql
select job_class_name, table_name, column_name, job_arguments from batched_background_migrations where status <> 3;

  
修改compose文件为14.3.6后up

6.  14.3.6->14.6.2  
    重复14.0.12 -> 14.3.6, 2-3步骤，sql结果为0行
7.  14.6.2->14.9.5  
    重复上一个版本升级过程（必须等待批量迁移完成)
8.  14.9.5->14.10.5  
    重复上一个版本升级过程（必须等待批量迁移完成)
9.  14.10.5->15.0.2  
    重复上一个版本升级过程（必须等待批量迁移完成)

##### 注: 查看批量迁移任务完成也可通过 Menu->Admin->Monitoring->Background Migrations 页面查看

#### gitlab更改root密码

gitlab-rails console
user = User.where(@root).first
user.password='12345678'
user.save!

#### gitlab后台设置用户为admin

gitlab-rails console
user = User.find_by(username: 'kurt')
user.admin = true
user.save!