# OpenSourceForHuaweiDemoGo

本项目作为 Go 版本的微服务开发示例，使用 [go-zero](https://go-zero.dev/) 作为微服务框架，数据库使用 GaussDB，项目结构按照 go-zero 最佳实践来布局。

数据库需要配置DSN：
1.在项目配置文件中 user/etc/user-api.yaml 配置；
2.使用环境变量 DSN 配置；

运行示例：
cd user
go run user.go

# 在华为云部署示例项目

参考：https://gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiWiki/blob/main/zh_CN/docs/cicd-pipeline.md

使用华为云 CodeArts 通过编译构建任务，将软件的源代码编译成镜像，并把镜像推送归档到容器镜像服务（SWR）中。

1. 打开 持续交付 -> 编译构建 -> 新建任务，配置构建任务

2. 选择模板，使用 系统模板 -> Go语言

构建步骤保留：构建环境配置、代码下载配置、Go语言构建、制作镜像并推送到SWR仓库。
构建环境配置：构建环境主机类型，选择 鲲鹏（ARM）服务器，执行主机选择 内置执行机。
Go语言构建：工具版本选择最新的Go版本，命令可以全部注释。

3. 完成配置，单击“保存并执行”。等待任务执行完毕，在镜像仓库会生成上述微服务的镜像。镜像内容可以通过 容器镜像服务 -> 我的镜像 进行查看。

4. 在CCE集群中创建自己的工作负载，并使用 ELB 服务对外暴露 RESTful 接口。


