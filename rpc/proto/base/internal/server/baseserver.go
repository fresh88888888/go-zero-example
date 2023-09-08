// Code generated by goctl. DO NOT EDIT.
// Source: base.proto

package server

import (
	"context"

	"umbrella.github.com/go-zero-example/rpc/proto/base/base"
	"umbrella.github.com/go-zero-example/rpc/proto/base/internal/logic"
	"umbrella.github.com/go-zero-example/rpc/proto/base/internal/svc"
)

type BaseServer struct {
	svcCtx *svc.ServiceContext
	base.UnimplementedBaseServer
}

func NewBaseServer(svcCtx *svc.ServiceContext) *BaseServer {
	return &BaseServer{
		svcCtx: svcCtx,
	}
}

func (s *BaseServer) Ping(ctx context.Context, in *base.Request) (*base.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}