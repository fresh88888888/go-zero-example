package logic

import (
	"context"

	"umbrella.github.com/go-zero-example/api/login/internal/svc"
	"umbrella.github.com/go-zero-example/api/login/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
