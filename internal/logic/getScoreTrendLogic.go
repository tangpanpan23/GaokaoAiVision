package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreTrendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreTrendLogic {
	return &GetScoreTrendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取分数线趋势
func (l *GetScoreTrendLogic) GetScoreTrend(in *score.GetScoreTrendReq) (*score.GetScoreTrendResp, error) {
	// todo: add your logic here and delete this line

	return &score.GetScoreTrendResp{}, nil
}
