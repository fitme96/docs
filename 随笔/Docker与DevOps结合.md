既然讲Docker与DevOps结合，我们先简单看看它们的定义

```
DevOps是一种软件开发流程和组织方式,强调开发团队和运维团队的沟通协作,以更高效地开发、部署和运维软件。它缩短开发与上线之间的周期,能快速响应业务需求的变化。
DevOps是一种文化、流程和组织方式,关注效率和质量。
```
```
Docker是一个开放源代码的开放平台软件，用于开发应用、交付（shipping）应用和运行应用。Docker允许用户将基础设施（Infrastructure）中的应用单独分割出来，形成更小的颗粒（容器），从而提高交付软件的速度
```
现在我们知道DevOps是一种文化，关注效率和质量。不知道也没关系，只要知道"CI/CD是实现DevOps的重要实践手段"，所以我们先搞CI/CD，加速交付，优化自己。

Docker的出现促进了CI/CD的实现，在Docker之前大家都是基于shell等脚本语言完成CI/CD，（不要杠说其他容器实现，没有标准化，小公司应该不会自己搞）虽然基本自动化，但是因为单个环境多次部署，版本回退，多个副本端口等问题导致的环境不一致问题，影响上线，所以我说一说我们对于Docker与CI/CD的结合。

Docker与CI/CD实现:
![[Pasted image 20230720143427.png]]

内部的一些说明：
- 因为公司产品是私有化部署，所以我们最终交付物是ISO，在这个过程中我们会生成Docker镜像并打包至ISO通过自动化安装ISO降低交付成本，提高部署成功率。
- CI/CD工具有GitLab、Jenkins、TravisCI等，今天我们使用Gitlab runner 来实现ci/cd，个人感觉对于小团队来说gitlab runner比较合适，在项目根目录编写.gitlab-ci.yml即可

使用的工具：
1. gitlab 、gitlab runner
3. 服务器 
4. 一个web项目
5. harbor仓库（用于存储docker镜像）
6. minIO (对象存储用于iso的保存)

以上工具说明：
- harbor是开源企业级Docker Registry，有很多功能， 如果不需要简单点就用docker官方registry
- minIO 对象存储，我们用于ISO的上传与下载。
- 服务器，如果runner负载很大，建议和gitlab分开部署


一些建议:
1. 为了避免多个runner管理混乱，建议由项目管理者注册一个group runner，开发者只需要在.gitlab-ci.yml中使用即可。
2. 镜像要分层（基础层、中间件层、应用层)，比如python 依赖不经常变动应该放入中间件层/或者再多一层）
3. 如果runner执行器是docker，那还有runner镜像Dockerfile也托管至项目根目录
4. ci/cd阶段多使用缓存，比如前端node_modules目录

也没写啥完整实现，我怕文章太长，所以就想着写了些，希望能帮助到你。