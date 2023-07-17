Deployment是更高阶的pod管理器，不应该使用底层的ReplicationController和ReplicaSet，当然Deployment创建了ReplicaSet，ReplicaSet管理pod。

### 为什么引入DeployMent，它的优点

### Deployment应用

-   创建deployment时增加  - -record(此选项会记录历史版本号，对于回退版本有很好的帮助)

### Deployment升级策略

1.  RollingUpdate 滚动升级。  默认升级策略
2.  Recreate 删除所有旧pod后创建新pod

### 回滚升级

-   kubectl rollout undo deployment test 回退到上一个版本
-   kubectl rollout history deployment test  查看历史版本
-   kubectl rollout undo deployment test - -to-revision=1   回退到指定版本
-   默认保留10个ReplicaSet,可通过revisionHistoryLimit做限制
-   默认10分钟不能完成滚动升级，视为失败，可以配置progressDeadlineSeconds设置deadline

### 暂停滚动升级/恢复滚动升级

-   kubectl rollout pause deployment test
-   kubectl rollout resume deployment test

通过暂停恢复可以完成金丝雀发布

### 通过就绪探针阻止全部滚动部署