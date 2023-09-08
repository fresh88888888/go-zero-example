package logic

import (
	"context"

	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/svc"
	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BookstoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookstoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BookstoreLogic {
	return &BookstoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookstoreLogic) Bookstore(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
