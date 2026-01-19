package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankToScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRankToScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankToScoreLogic {
	return &GetRankToScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取位次对应的分数
func (l *GetRankToScoreLogic) GetRankToScore(in *score.GetRankToScoreReq) (*score.GetRankToScoreResp, error) {
	// todo: add your logic here and delete this line

	return &score.GetRankToScoreResp{}, nil
}
