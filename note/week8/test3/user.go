package main

import (
	"flag"
	"fmt"

	"test3/github.com/example/user"
	"test3/internal/config"
	userclassserviceServer "test3/internal/server/userclassservice"
	userroleserviceServer "test3/internal/server/userroleservice"
	userserviceServer "test3/internal/server/userservice"
	"test3/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
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
		user.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		user.RegisterUserRoleServiceServer(grpcServer, userroleserviceServer.NewUserRoleServiceServer(ctx))
		user.RegisterUserClassServiceServer(grpcServer, userclassserviceServer.NewUserClassServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
