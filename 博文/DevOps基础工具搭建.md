

# 解决的问题
- 减少开发人员处理环境问题
- 避免测试人员自建测试环境
- 加快转测/发版

# 最终实现的效果
开发者提交MR -> 合并触发构建 -> 自动完成更新


# 需要了解的前置知识
- Docker
- Dockerfile
- compose 
- git

# 需要搭建的基础工具
- gitlab
- gitlab runner
- harbor

## gitlab
```yaml
version: '3.6'
services:
  gitlab:
    image: 'gitlab/gitlab-ce:15.0.2-ce.0'
    hostname: 'git.dd.com'
    restart: always
    ports:
      - '127.0.0.1:8929:8929'
      - '22:22'
    volumes:
      - '/data/gitlab/config:/etc/gitlab'
      - '/data/gitlab/logs:/var/log/gitlab'
      - '/data/gitlab/data:/var/opt/gitlab'
```

## gitlab runner
```bash
 docker run -d --name runner --restart always \
     -v /data/runner/config:/etc/gitlab-runner \
     -v /var/run/docker.sock:/var/run/docker.sock \
     gitlab/gitlab-runner:latest

```

## harbor


