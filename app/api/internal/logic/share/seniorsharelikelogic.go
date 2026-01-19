package share

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeniorShareLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeniorShareLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeniorShareLikeLogic {
	return &SeniorShareLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeniorShareLikeLogic) SeniorShareLike(req *types.SeniorShareLikeReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
