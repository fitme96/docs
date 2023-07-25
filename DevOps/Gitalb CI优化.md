runner采用docker执行器

## 后端

### go

-   runner增加volume到GOPATH/pkg/mod,这样可以采用本地goproxy缓存，不用每次走下载
-   在多套环境中，为了增加缓存命中率可以采用cache keys方式
-   Dockerfile禁止使用go builder，减少启动多个容器

image: hub.bugfeel.net:8443/custom/golang-1.17.5-modsecurity-dockercli:latest
stages:
  - test
  - build
  - deploy

.test_template:
  script:
    - export GOPATH="$CI_PROJECT_DIR/cache"
    - go test ./pkg/modsecurity 
    - go test ./pkg/middlewares/basewafrule 
    - ls $CI_PROJECT_DIR/cache
  tags:
    - docker-runner

.build_template:
  cache:
    key: ${CI_COMMIT_REF_SLUG}
    paths:
      - cache
  script:
    - export GOPATH="$CI_PROJECT_DIR/cache"
    - go build -ldflags="-s -w" -o dist/traefik ./cmd/traefik
    - docker login hub.bugfeel.net:8443 -u $DOCKER_USER -p $DOCKER_PASSWD
    - docker build -t $DOCKER_REGISTRY/$NAMESPACE/$IMAGE_NAME:`date +%Y-%m-%d`-${CI_COMMIT_SHORT_SHA} -f Dockerfile-app .
    - docker tag $DOCKER_REGISTRY/$NAMESPACE/$IMAGE_NAME:`date +%Y-%m-%d`-${CI_COMMIT_SHORT_SHA} $DOCKER_REGISTRY/$NAMESPACE/$IMAGE_NAME:latest
    - docker push $DOCKER_REGISTRY/$NAMESPACE/$IMAGE_NAME:`date +%Y-%m-%d`-${CI_COMMIT_SHORT_SHA}
    - docker push $DOCKER_REGISTRY/$NAMESPACE/$IMAGE_NAME:latest
  tags:
    - docker-runner


variables:
  IMAGE_NAME: "traefik"

# 开发环境单测
dev-test:
  stage: test
  extends: .test_template
  only:
    - dev
# 开发环境构建镜像
dev-build:
  stage: build
  variables:
    NAMESPACE: "yulei-dev"
  extends: .build_template
  only:
    - dev
# master单测
master-test:
  stage: test
  extends: .test_template
  only:
    - tags
# master构建镜像
master-build:
  stage: build
  variables:
    NAMESPACE: "yulei-master"
  extends: .build_template
  only:
    - tags

## 前端

### node

-   做cache(比如node_modules)时增加key:$CI_JOB_STAGE-$CI_COMMIT_REF_SLUG这样可以增加缓存命中率
-   增加npm私服可以提升下载依赖速度

## 效果

-   后端从1.30s-2.10s 提升到30s-50s
-   前端从4m 提升到2m

## 思考

-   对于后端来说提升空间有限，前端提升空间还是很大的，node构建速度的提升需要深入研究一下.