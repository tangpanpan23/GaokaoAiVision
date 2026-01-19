package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeniorSharesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSeniorSharesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeniorSharesLogic {
	return &GetSeniorSharesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学长分享列表
func (l *GetSeniorSharesLogic) GetSeniorShares(in *user.GetSeniorSharesReq) (*user.GetSeniorSharesResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetSeniorSharesResp{}, nil
}
