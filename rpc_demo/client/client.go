package main

import (
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"umbrella.github.com/go-zero-example/rpc_demo/client/greet"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/greet-client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	defer conn.Conn().Close()
	resp, err := greet.NewGreetClient(conn.Conn()).Ping(context.Background(), &greet.Request{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp)
}
