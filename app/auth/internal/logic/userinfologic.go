package logic

import (
	"context"
	"errors"
	"strings"

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

func (l *UserinfoLogic) Userinfo(req *types.UserInfoRequest) (*types.UserInfoResponse, error) {
	authHeader := req.AccessToken
	if authHeader == "" {
		return nil, errors.New("no access token")
	}

	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		return nil, errors.New("token is not Bearer token")
	}

	claims, err := casdoorsdk.ParseJwtToken(token[1])
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
