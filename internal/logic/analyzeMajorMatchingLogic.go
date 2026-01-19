package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeMajorMatchingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeMajorMatchingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeMajorMatchingLogic {
	return &AnalyzeMajorMatchingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 专业匹配度分析
func (l *AnalyzeMajorMatchingLogic) AnalyzeMajorMatching(in *ai.AnalyzeMajorMatchingReq) (*ai.AnalyzeMajorMatchingResp, error) {
	// todo: add your logic here and delete this line

	return &ai.AnalyzeMajorMatchingResp{}, nil
}
