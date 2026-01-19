package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateVolunteerSuggestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateVolunteerSuggestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateVolunteerSuggestionLogic {
	return &GenerateVolunteerSuggestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成志愿推荐
func (l *GenerateVolunteerSuggestionLogic) GenerateVolunteerSuggestion(in *ai.GenerateVolunteerSuggestionReq) (*ai.GenerateVolunteerSuggestionResp, error) {
	// todo: add your logic here and delete this line

	return &ai.GenerateVolunteerSuggestionResp{}, nil
}
