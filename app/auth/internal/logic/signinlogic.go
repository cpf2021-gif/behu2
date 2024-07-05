package logic

import (
	"context"
	"errors"

	"behu2/app/auth/internal/svc"
	"behu2/app/auth/internal/types"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/oauth2"
)

type SigninLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigninLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigninLogic {
	return &SigninLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SigninLogic) Signin(req *types.SignInRequest) (*types.SignInResponse, *oauth2.Token, error) {
	token, err := casdoorsdk.GetOAuthToken(req.Code, req.State)
	if err != nil {
		return nil, nil, errors.New("failed to get OAuth token")
	}

	return &types.SignInResponse{
		Status: "ok",
	}, token, nil
}
