package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AIChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAIChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AIChatLogic {
	return &AIChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AIChatLogic) AIChat(req *types.AIChatRequest) (resp *types.AIChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
