package main

import (
	"flag"
	"time"

	"database-service/internal/config"
	"database-service/internal/server"
	"database-service/internal/svc"
	"database-service/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"git.ickey.com.cn/infra/utils/grpc/interceptor"
	"git.ickey.com.cn/infra/utils/initialize"
	"git.ickey.com.cn/infra/utils/log"
	"git.ickey.com.cn/infra/zero-contrib/consul"
)

var (
	configFile = flag.String("f", "etc/app_local.yaml", "the config file")
	serviceTag = flag.String("t", "", "service tag for dyeing")
)
var (
	branch       = ""
	commit       = ""
	goctlVersion = "v2.1.4"
)

func main() {
	flag.Parse()

	initialize.Init(*configFile)
	log.Infof("app branch: %s", branch)
	log.Infof("app commit: %s", commit)
	log.Infof("goctl version: %s", goctlVersion)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewDatabaseServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterDatabaseServiceServer(grpcServer, srv)
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 注册中间件
	s.AddUnaryInterceptors(interceptor.GetUnaryServerInterceptor()...)

	// 注册服务tag
	if serviceTag != nil && *serviceTag != "" {
		c.Consul.Tag = append(c.Consul.Tag, *serviceTag)
	}

	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		log.Fatalf("register consul service failed: %v", err)
		return
	}

	proc.SetTimeToForceQuit(time.Duration(c.TimeBeforeQuit) * time.Second)
	//启动服务
	log.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
	<-make(chan struct{})
}
