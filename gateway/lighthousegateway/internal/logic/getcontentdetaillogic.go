package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentDetailLogic {
	return &GetContentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentDetailLogic) GetContentDetail(req *types.EmptyRequest) (resp *types.GetContentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
