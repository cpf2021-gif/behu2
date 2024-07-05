package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetCookieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCookieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCookieLogic {
	return &SetCookieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCookieLogic) SetCookie(req *types.SetCookieRequest) (resp *types.SetCookieResponse, refreshToken string, expiry time.Time, err error) {
	authHeader := req.RefreshToken
	if authHeader == "" {
		err = errors.New("authHeader is empty")
		return
	}

	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		return nil, "", time.Now(), errors.New("token is not Bearer token")
	}

	newTokenPair, err := casdoorsdk.RefreshOAuthToken(token[1])
	if err != nil {
		return nil, "", time.Now(), errors.New("refresh token failed")
	}

	return &types.SetCookieResponse{
		Status:      "ok",
		AccessToken: newTokenPair.AccessToken,
	}, newTokenPair.RefreshToken, newTokenPair.Expiry, nil
}
