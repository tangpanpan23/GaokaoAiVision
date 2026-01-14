package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentLogic) GetContent(req *types.GetContentRequest) (resp *types.GetContentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
