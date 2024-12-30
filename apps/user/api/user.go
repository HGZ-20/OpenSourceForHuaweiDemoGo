package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/config"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/handler"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 读取环境变量并更换为 Target
	target := os.Getenv("Target")
	if target == "" {
		panic("Target environment variable is not set")
	}
	// 读取环境变量并解析为 []string
	endpointsStr := os.Getenv("ENDPOINTS")
	if endpointsStr == "" {
		panic("ENDPOINTS environment variable is not set")
	}
	c.UserRpc.Endpoints = strings.Split(endpointsStr, ",")
	c.UserRpc.Target = target

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
