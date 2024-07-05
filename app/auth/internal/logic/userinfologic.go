package logic

import (
	"context"
	"errors"

	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo() (*types.UserInfoResponse, error) {
	cookie := l.ctx.Value("token")
	var token string
	if cookieStr, ok := cookie.(string); ok {
		token = cookieStr
	} else {
		return nil, errors.New("invalid token")
	}

	claims, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		return nil, errors.New("failed to parse token")
	}

	return &types.UserInfoResponse{
		Status: "ok",
		Data: types.UserInfo{
			DisplayName: claims.DisplayName,
			Avatar:      claims.Avatar,
		},
	}, nil
}
