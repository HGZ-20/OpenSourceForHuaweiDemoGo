# OpenSourceForHuaweiDemoGo

本项目作为 Go 版本的微服务开发示例，使用 [go-zero](https://go-zero.dev/) 作为微服务框架，数据库使用华为云 GaussDB，项目结构按照 go-zero 最佳实践来布局。

git clone https://gitcode.com/yangjiaxin/OpenSourceForHuaweiDemoGo.git

cd OpenSourceForHuaweiDemoGo

1.初始化数据库：

在华为云 GaussDB 数据库终端执行 apps/user/model/user.sql 语句，创建用户表。

2.修改rpc服务下的数据库需要配置 DataSource：

(1)在项目配置文件中 apps/user/rpc/etc/user-rpc.yaml 配置； 

(2)或者使用环境变量 DataSource 配置；

3.部署Nacos，分别在 rpc 和 api 中配置服务注册中心地址：<br>
(1)在项目配置文件中 apps/user/rpc/etc/user-rpc.yaml 修改 Nacos 配置；<br>
(2)在项目配置文件中 apps/user/api/etc/user-api.yaml 中的 UserRpc 修改 Target 中 nacos 的地址。(可选) <br>

4.运行示例： 

先启动rpc服务，再启动api服务<br>
`go run apps/user/rpc/user.go -f apps/user/rpc/etc/user-rpc.yaml`

`go run apps/user/api/user.go -f apps/user/api/etc/user-api.yaml`
或者<br>
参考 Makefile 运行<br>

# 在华为云部署示例项目

使用华为云 CodeArts 通过编译构建任务，将软件的源代码编译成镜像，并把镜像推送归档到容器镜像服务（SWR）中。

参考Wiki：https://gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiWiki/blob/main/zh_CN/docs/go/cicd-pipeline.md

![Arch](arch.png)


