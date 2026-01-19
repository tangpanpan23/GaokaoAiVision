package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeSeniorShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeSeniorShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeSeniorShareLogic {
	return &LikeSeniorShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞/取消点赞分享
func (l *LikeSeniorShareLogic) LikeSeniorShare(in *user.LikeSeniorShareReq) (*user.LikeSeniorShareResp, error) {
	// todo: add your logic here and delete this line

	return &user.LikeSeniorShareResp{}, nil
}
