package config

import (
	"github.com/tal-tech/go-zero/zrpc"

)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache   cache.CacheConf
}

