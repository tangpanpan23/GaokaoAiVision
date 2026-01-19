package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeUserProfileLogic {
	return &AnalyzeUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分析用户画像
func (l *AnalyzeUserProfileLogic) AnalyzeUserProfile(in *ai.AnalyzeUserProfileReq) (*ai.AnalyzeUserProfileResp, error) {
	// todo: add your logic here and delete this line

	return &ai.AnalyzeUserProfileResp{}, nil
}
