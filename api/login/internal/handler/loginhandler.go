package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"umbrella.github.com/go-zero-example/api/login/internal/logic"
	"umbrella.github.com/go-zero-example/api/login/internal/svc"
	"umbrella.github.com/go-zero-example/api/login/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
