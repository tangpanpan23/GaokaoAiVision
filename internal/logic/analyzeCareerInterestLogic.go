package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeCareerInterestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeCareerInterestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeCareerInterestLogic {
	return &AnalyzeCareerInterestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 职业兴趣分析
func (l *AnalyzeCareerInterestLogic) AnalyzeCareerInterest(in *ai.AnalyzeCareerInterestReq) (*ai.AnalyzeCareerInterestResp, error) {
	// todo: add your logic here and delete this line

	return &ai.AnalyzeCareerInterestResp{}, nil
}
