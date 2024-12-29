package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpc zrpc.RpcClientConf `json:",env=UserRpc"`
	//Target   string             `json:",env=Target"`
	//Endpoint string             `json:",env=Endpoint"`
}

// RpcClientConfig 用于包装zrpc.RpcClientConf并绑定环境变量
type RpcClientConfig struct {
	//Etcd          string        `json:",env=Etcd"`
	Endpoint string `json:",env=Endpoint"`
	Target   string `json:",env=Target"`
	Timeout  int64  `json:",env=Timeout"`
}

// ToRpcClientConf 将包装后的配置转化为 zrpc.RpcClientConf
func (c *RpcClientConfig) ToRpcClientConf() zrpc.RpcClientConf {
	var endpoints []string
	endpoints = append(endpoints, c.Endpoint)
	return zrpc.RpcClientConf{
		Endpoints: endpoints,
		Target:    c.Target,
		Timeout:   c.Timeout,
	}
}
