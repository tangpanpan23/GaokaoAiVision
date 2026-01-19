package volunteer

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreQueryLogic {
	return &ScoreQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreQueryLogic) ScoreQuery(req *types.ScoreQueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
