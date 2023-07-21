### RBAC鉴权开启

kube-apiserver --authorization-mode=Example,RBAC --<其他选项> --<其他选项>

RBAC API 声明了四种 Kubernetes 对象：Role、ClusterRole、RoleBinding 和 ClusterRoleBinding。

-   RBAC用来保护集群安全性，防止恶意API调用（比如修改集群数据，破坏集群),RBAC授权插件将用户角色作为决定用户是否执行操作的关键因素。主体(可以是一个人，一个ServiceAccount,或者一组用户或ServiceAccoutn)和一个或多个角色相关联。

[  
](https://www.notion.so/40a55991271246d999fd8129cbff859d)

### ServiceAccount

ServiceAccount是对于pod而言，每个namespace都有一个default ServiceAccount，在声明文件中没有显式添加SerivceAccount会默认使用default

ServiceAccount作用:

1.  增加imagePullSecrets字段可以避免为每个pod增加拉取镜像密钥
2.  限制pod可挂载的密钥（开启限制需要在ServiceAccount增加注解:kubernetes.io/enforce-mountable-secrets="true".

### UserAccount

UserAccount是对于真实用户而言比如kubectl执行

## 生产环境K8S思考

-   对于不同的开发者创建各自的家目录配置最小的权限
-   拉取镜像添加到ServiceAccount