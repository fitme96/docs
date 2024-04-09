# runner采用docker执行器



### go

-   runner增加volume到GOPATH/pkg/mod,这样可以采用本地goproxy缓存，不用每次走下载
-   在多套环境中，为了增加缓存命中率可以采用cache keys方式
-   Dockerfile禁止使用go builder，减少启动多个容器


### python

```yaml



```

### node
```yaml
# 前端构建模板
.web_build:
  image: hub.bugfeel.net:8443/ad-train/python:3.9.16-node16.20.2-ssh
  stage: build
  cache:
    key: ${CI_COMMIT_REF_SLUG}-${pro_module}
    paths:
      - ${pro_module}/node_modules
  script:
    - pnpm config set registry https://registry.npmmirror.com
    - cd ${pro_module}
    - pnpm install
    - pnpm run build
    - docker login hub.bugfeel.net:8443 -u $DOCKER_USER -p $DOCKER_PASSWD
    - docker buildx build --platform linux/amd64 --build-arg DIR=${dist} -t hub.bugfeel.net:8443/ad-train/${pro_module}:`date +%Y-%m-%d`-${CI_COMMIT_SHORT_SHA}  -t  hub.bugfeel.net:8443/ad-train/${pro_module}:latest -f ../Dockerfile-web --push .
  tags:
    - ad-train

```

-   做cache(比如node_modules)时增加key:$CI_JOB_STAGE-$CI_COMMIT_REF_SLUG这样可以增加缓存命中率
-   增加npm私服可以提升下载依赖速度

