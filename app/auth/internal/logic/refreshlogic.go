package logic

import (
	"context"
	"errors"
	"time"

	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh() (resp *types.RefreshResponse, refreshToken string, expiry time.Time, err error) {
	token := l.ctx.Value("refresh_token")
	if tokenstr, ok := token.(string); ok {
		refreshToken = tokenstr
	} else {
		err = errors.New("invalid refresh token")
		return
	}

	newTokenPair, err := casdoorsdk.RefreshOAuthToken(refreshToken)
	if err != nil {
		err = errors.New("refresh token failed")
		return
	}

	resp = &types.RefreshResponse{
		Status:      "ok",
		AccessToken: newTokenPair.AccessToken,
	}

	expiry = newTokenPair.Expiry
	return
}
