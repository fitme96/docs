# runner采用docker执行器

## 后端

### go

-   runner增加volume到GOPATH/pkg/mod,这样可以采用本地goproxy缓存，不用每次走下载
-   在多套环境中，为了增加缓存命中率可以采用cache keys方式
-   Dockerfile禁止使用go builder，减少启动多个容器

### node

-   做cache(比如node_modules)时增加key:$CI_JOB_STAGE-$CI_COMMIT_REF_SLUG这样可以增加缓存命中率
-   增加npm私服可以提升下载依赖速度

## 效果

-   后端从1.30s-2.10s 提升到30s-50s
-   前端从4m 提升到2m

## 思考

-   对于后端来说提升空间有限，前端提升空间还是很大的，node构建速度的提升需要深入研究一下.
