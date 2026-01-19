package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetCollegeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetCollegeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetCollegeInfoLogic {
	return &BatchGetCollegeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取学校信息
func (l *BatchGetCollegeInfoLogic) BatchGetCollegeInfo(in *score.BatchGetCollegeInfoReq) (*score.BatchGetCollegeInfoResp, error) {
	// todo: add your logic here and delete this line

	return &score.BatchGetCollegeInfoResp{}, nil
}
