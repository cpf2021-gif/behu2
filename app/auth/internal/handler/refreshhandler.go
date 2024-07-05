package handler

import (
	"context"
	"errors"
	"net/http"

	"behu2/app/auth/internal/logic"
	"behu2/app/auth/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func refreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get cookie
		refreshToken, err := r.Cookie("refresh_token")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("no refresh token"))
		}

		// set refresh token to context
		ctx := context.WithValue(r.Context(), "refresh_token", refreshToken.Value)

		l := logic.NewRefreshLogic(ctx, svcCtx)
		resp, token, expiry, err := l.Refresh()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// set new refresh token to cookie
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
