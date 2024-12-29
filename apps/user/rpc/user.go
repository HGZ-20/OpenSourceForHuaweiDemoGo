package main

import (
	"flag"
	"fmt"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/config"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/server"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// register service to nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.IP, c.Nacos.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         c.Nacos.NamespaceId,
		TimeoutMs:           c.Nacos.TimeoutMs,
		NotLoadCacheAtStart: c.Nacos.NotLoadCacheAtStart,
		LogDir:              c.Nacos.LogDir,
		CacheDir:            c.Nacos.CacheDir,
		LogLevel:            c.Nacos.LogLevel,
	}

	opts := nacos.NewNacosConfig(c.RpcServerConf.Name, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)

	//gw := gateway.MustNewServer(c.Gateway)
	//group := service.NewServiceGroup()
	//group.Add(s)
	//group.Add(gw)
	//
	//defer group.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	//fmt.Printf("Starting gateway at %s:%d...\n", c.Gateway.Host, c.Gateway.Port)
	s.Start()
}
