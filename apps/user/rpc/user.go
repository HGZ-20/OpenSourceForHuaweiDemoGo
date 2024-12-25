package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/config"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/server"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		*constant.NewServerConfig(c.IP, c.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         c.NamespaceId,
		TimeoutMs:           c.TimeoutMs,
		NotLoadCacheAtStart: c.NotLoadCacheAtStart,
		LogDir:              c.LogDir,
		CacheDir:            c.CacheDir,
		LogLevel:            c.LogLevel,
	}

	opts := nacos.NewNacosConfig(c.RpcServerConf.Name, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
