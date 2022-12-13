package config

import (
	"git.ickey.com.cn/infra/zero-contrib/consul"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Consul         consul.Conf
	DemoRpc        zrpc.RpcClientConf
	TimeBeforeQuit int
}
