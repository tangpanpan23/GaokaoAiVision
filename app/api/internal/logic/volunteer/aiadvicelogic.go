package volunteer

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AIAdviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAIAdviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AIAdviceLogic {
	return &AIAdviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AIAdviceLogic) AIAdvice(req *types.AIAdviceReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
