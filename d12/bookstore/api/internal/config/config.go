package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	Add zrpc.RpcClientConf
	Check zrpc.RpcClientConf
}
