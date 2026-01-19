package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelSpiderTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelSpiderTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelSpiderTaskLogic {
	return &CancelSpiderTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消爬虫任务
func (l *CancelSpiderTaskLogic) CancelSpiderTask(in *spider.CancelSpiderTaskReq) (*spider.CancelSpiderTaskResp, error) {
	// todo: add your logic here and delete this line

	return &spider.CancelSpiderTaskResp{}, nil
}
