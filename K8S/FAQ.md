问题：

x509: certificate signed by unknown authority 

解决:

镜像没有 ca-certificates 根证书，导致无法识别外部 https 携带的数字证书
apt-get -qq update  apt-get -qq install -y --no-install-recommends ca-certificates curl


### 删除namespace后一直处于Terminating

1.  kubectl get ns bar -o json > tmp.json
2.  编辑tmp.json,删除finalizers 字段的值
3.  kubectl proxy 启动proxy
4.  curl -k -H "Content-Type: application/json" -X PUT --data-binary [@tmp.json](/tmp.json) [http://127.0.0.1:8001/api/v1/namespaces/foo/finalize](http://127.0.0.1:8001/api/v1/namespaces/foo/finalize)
5.  完成Terminating ns删除

### V1.20.0 版本出现selfLink was empty

-   详细日志：

E1210 14:42:01.500487 1 controller.go:1004] provision "default/pvc1" class "managed-nfs-storage": unexpected error getting claim reference: selfLink was empty, can't make reference
E1210 14:42:01.500502 1 controller.go:1004] provision "default/test-claim" class "managed-nfs-storage": unexpected error getting claim reference: selfLink was empty, can't make reference

-   解决：

在apiserver启动命令 增加 --feature-gates=RemoveSelfLink=false行

-   参考

[https://stackoverflow.com/questions/65376314/kubernetes-nfs-provider-selflink-was-empty](https://stackoverflow.com/questions/65376314/kubernetes-nfs-provider-selflink-was-empty)

[https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner/issues/25](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner/issues/25)

### ingress-nginx 泛域名匹配

[nginx.ingress.kubernetes.io/server-alias](http://nginx.ingress.kubernetes.io/server-alias:)

eg:

metadata:
  name: hicloud-frontend
  namespace: user-center
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/use-regex: "true"
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/rewrite-target: /$1
    nginx.ingress.kubernetes.io/server-alias: '~^(.+)\.hicloud\.hive-intel\.qa$'

### ingress-nginx 多域名

eg:

- host: "api.hive-intel.qa"
    http:
      paths:
      - path: "/user-center(/|$)(.*)"
        pathType: Prefix
        backend:
          service:
            name: user-center
            port:
              number: 8080
  - host: "apiqa.hive-intel.com"
    http:
      paths:
      - path: "/user-center(/|$)(.*)"
        pathType: Prefix
        backend:
          service:
            name: user-center
            port:
              number: 8080