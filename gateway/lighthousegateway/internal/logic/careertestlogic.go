package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CareerTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCareerTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CareerTestLogic {
	return &CareerTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CareerTestLogic) CareerTest(req *types.CareerTestRequest) (resp *types.CareerTestResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
