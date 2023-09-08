package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/logic"
	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/svc"
	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/types"
)

func BookstoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBookstoreLogic(r.Context(), svcCtx)
		resp, err := l.Bookstore(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
