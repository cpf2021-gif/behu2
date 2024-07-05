package handler

import (
	"net/http"

	"behu2/app/auth/internal/logic"
	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func setCookieHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetCookieRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSetCookieLogic(r.Context(), svcCtx)
		resp, token, expiry, err := l.SetCookie(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			cookie := http.Cookie{
				Name:  "refresh_token",
				Value: token,

				Path:    "/",
				Domain:  "localhost",
				Expires: expiry,

				HttpOnly: true,
				Secure:   false,
			}

			http.SetCookie(w, &cookie)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
