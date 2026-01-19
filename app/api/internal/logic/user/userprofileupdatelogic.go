package user

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserProfileUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserProfileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserProfileUpdateLogic {
	return &UserProfileUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserProfileUpdateLogic) UserProfileUpdate(req *types.UserProfileUpdateReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
