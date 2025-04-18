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

	// 读取环境变量并替换 Target 中的占位符
	nacosIP := os.Getenv("NACOS_IP")
	if nacosIP == "" {
		fmt.Errorf("NACOS_IP environment variable is not set")
	}
	// 读取环境变量并解析为 []string
	endpointsStr := os.Getenv("ENDPOINTS")
	if endpointsStr == "" {
		fmt.Errorf("ENDPOINTS environment variable is not set")
	}
	c.UserRpc.Endpoints = strings.Split(endpointsStr, ",")
	c.UserRpc.Target = strings.Replace(c.UserRpc.Target, "{NACOS_IP}", nacosIP, 1)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
