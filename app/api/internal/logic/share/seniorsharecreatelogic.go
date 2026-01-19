package share

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeniorShareCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeniorShareCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeniorShareCreateLogic {
	return &SeniorShareCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeniorShareCreateLogic) SeniorShareCreate(req *types.SeniorShareCreateReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
