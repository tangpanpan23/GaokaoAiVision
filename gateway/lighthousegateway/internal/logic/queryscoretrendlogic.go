package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryScoreTrendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryScoreTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryScoreTrendLogic {
	return &QueryScoreTrendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryScoreTrendLogic) QueryScoreTrend(req *types.QueryScoreRequest) (resp *types.QueryScoreResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
