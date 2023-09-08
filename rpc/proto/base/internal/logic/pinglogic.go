package logic

import (
	"context"

	"umbrella.github.com/go-zero-example/rpc/proto/base/base"
	"umbrella.github.com/go-zero-example/rpc/proto/base/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *base.Request) (*base.Response, error) {
	// todo: add your logic here and delete this line

	return &base.Response{}, nil
}
