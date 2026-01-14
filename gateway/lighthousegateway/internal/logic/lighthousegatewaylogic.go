package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LighthousegatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLighthousegatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LighthousegatewayLogic {
	return &LighthousegatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LighthousegatewayLogic) Lighthousegateway(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
