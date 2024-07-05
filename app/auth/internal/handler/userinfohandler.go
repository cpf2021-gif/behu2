package handler

import (
	"context"
	"errors"
	"net/http"

	"behu2/app/auth/internal/logic"
	"behu2/app/auth/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("no token"))
		}

		ctx := context.WithValue(r.Context(), "token", token.Value)

		l := logic.NewUserinfoLogic(ctx, svcCtx)
		resp, err := l.Userinfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
