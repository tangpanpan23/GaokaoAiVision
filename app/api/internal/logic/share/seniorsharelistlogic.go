package share

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeniorShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeniorShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeniorShareListLogic {
	return &SeniorShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeniorShareListLogic) SeniorShareList(req *types.SeniorShareListReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
