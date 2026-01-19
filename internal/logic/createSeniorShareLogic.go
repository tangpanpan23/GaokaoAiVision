package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeniorShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeniorShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeniorShareLogic {
	return &CreateSeniorShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建学长分享
func (l *CreateSeniorShareLogic) CreateSeniorShare(in *user.CreateSeniorShareReq) (*user.CreateSeniorShareResp, error) {
	// todo: add your logic here and delete this line

	return &user.CreateSeniorShareResp{}, nil
}
