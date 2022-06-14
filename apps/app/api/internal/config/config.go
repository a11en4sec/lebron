package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	OrderRPC   zrpc.RpcClientConf
	ProductRPC zrpc.RpcClientConf
}
