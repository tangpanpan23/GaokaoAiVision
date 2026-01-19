package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetMajorInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetMajorInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetMajorInfoLogic {
	return &BatchGetMajorInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取专业信息
func (l *BatchGetMajorInfoLogic) BatchGetMajorInfo(in *score.BatchGetMajorInfoReq) (*score.BatchGetMajorInfoResp, error) {
	// todo: add your logic here and delete this line

	return &score.BatchGetMajorInfoResp{}, nil
}
