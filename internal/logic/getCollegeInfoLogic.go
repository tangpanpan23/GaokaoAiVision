package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollegeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollegeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollegeInfoLogic {
	return &GetCollegeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学校信息
func (l *GetCollegeInfoLogic) GetCollegeInfo(in *score.GetCollegeInfoReq) (*score.GetCollegeInfoResp, error) {
	// todo: add your logic here and delete this line

	return &score.GetCollegeInfoResp{}, nil
}
