package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpiderTargetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpiderTargetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpiderTargetsLogic {
	return &GetSpiderTargetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取爬虫目标列表
func (l *GetSpiderTargetsLogic) GetSpiderTargets(in *spider.GetSpiderTargetsReq) (*spider.GetSpiderTargetsResp, error) {
	// todo: add your logic here and delete this line

	return &spider.GetSpiderTargetsResp{}, nil
}
