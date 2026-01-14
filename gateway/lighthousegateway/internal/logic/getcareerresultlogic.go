package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareerResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCareerResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareerResultLogic {
	return &GetCareerResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCareerResultLogic) GetCareerResult(req *types.EmptyRequest) (resp *types.CareerTestResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
