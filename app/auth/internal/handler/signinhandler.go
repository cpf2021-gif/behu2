package handler

import (
	"net/http"

	"behu2/app/auth/internal/logic"
	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func signinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSigninLogic(r.Context(), svcCtx)

		resp, err := l.Signin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
