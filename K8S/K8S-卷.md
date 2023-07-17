K8S中可通过挂载卷来完成数据持久化  
  

卷的类型  
  
●emptyDir # 临时卷，随着pod销毁  
●hostPath # 可以映射到主机特定目录，就相当与docker中的-v参数，和主机耦合高  
●gitRepo # git仓库内容初始化的卷  
●nfs # 挂载nfs  
●cephfs,rbd(ceph块设备）# 其他网络类型存储  
●configMap 、secret # 特殊类型的卷  
●persistentVolumeClaim # 持久存储类型区别于数据持久化  
  

卷的示例  
  

PV/PVC  
  
k8s中引入了pv\pvc可以对卷和pod的接耦合，屏蔽基础设施细节，开发人员不需要关注存储是如何实现的，只需要在pvc中做申请使用即可，集群管理员提供pv。这也是符合k8s理念的  
  

动态卷  
  
●在K8S中大部分都是无状态服务，对于持久化存储需求不大，但是要部署有状态应用，比如MongoDB Sharding  会有很多问题  
  
1共享存储肯定不行，因为它们有自己的数据比如一个shard会有副本集共享存储不能启动多个mongodb进程了  
2HostPath也不行，一个node跑多个副本也会冲突  
3静态pv，pvc也不行工程大，和1会遇到相同的问题  
  
动态卷，是将一个网络存储作为一个StorageClass类，来动态的创建pv,pvc并进行绑定，可以实现动态的存储生成与持久化保存。  

若有收获，就点个赞吧