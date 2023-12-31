package main

import (
	"flag"
	"fmt"

	"umbrella.github.com/go-zero-example/bookstore/rpc/add/add"
	"umbrella.github.com/go-zero-example/bookstore/rpc/add/internal/config"
	"umbrella.github.com/go-zero-example/bookstore/rpc/add/internal/server"
	"umbrella.github.com/go-zero-example/bookstore/rpc/add/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/add.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAddServer(grpcServer, server.NewAddServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
