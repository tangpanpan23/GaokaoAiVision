package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAdmissionScoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAdmissionScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdmissionScoresLogic {
	return &QueryAdmissionScoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询录取分数线
func (l *QueryAdmissionScoresLogic) QueryAdmissionScores(in *score.QueryAdmissionScoresReq) (*score.QueryAdmissionScoresResp, error) {
	// todo: add your logic here and delete this line

	return &score.QueryAdmissionScoresResp{}, nil
}
