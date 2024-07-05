package handler

import (
	"net/http"
	"time"

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

		resp, token, err := l.Signin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			accessTokenCookie := http.Cookie{
				Name:  "token",
				Value: token.AccessToken,

				Path:    "/",
				Domain:  "localhost",
				Expires: token.Expiry,

				HttpOnly: true,
				Secure:   false,
			}
			refreshTokenCookie := http.Cookie{
				Name:  "refresh_token",
				Value: token.RefreshToken,

				Path:    "/",
				Domain:  "localhost",
				Expires: time.Now().Add(6 * 24 * time.Hour),

				HttpOnly: true,
				Secure:   false,
			}
			http.SetCookie(w, &accessTokenCookie)
			http.SetCookie(w, &refreshTokenCookie)

			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
