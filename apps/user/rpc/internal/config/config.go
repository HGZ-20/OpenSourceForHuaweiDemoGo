package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Nacos struct {
		IP                  string
		Port                uint64
		NamespaceId         string
		NotLoadCacheAtStart bool
		LogLevel            string
		LogDir              string
		CacheDir            string
		TimeoutMs           uint64
	}

	GaussDB struct {
		DataSource string
	}
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}

	Salt string
}
