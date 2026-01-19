package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/score/rpc/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMajorInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMajorInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMajorInfoLogic {
	return &GetMajorInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取专业信息
func (l *GetMajorInfoLogic) GetMajorInfo(in *score.GetMajorInfoReq) (*score.GetMajorInfoResp, error) {
	// todo: add your logic here and delete this line

	return &score.GetMajorInfoResp{}, nil
}
