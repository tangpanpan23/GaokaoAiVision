package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpiderResultsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpiderResultsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpiderResultsLogic {
	return &GetSpiderResultsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取爬取结果
func (l *GetSpiderResultsLogic) GetSpiderResults(in *spider.GetSpiderResultsReq) (*spider.GetSpiderResultsResp, error) {
	// todo: add your logic here and delete this line

	return &spider.GetSpiderResultsResp{}, nil
}
