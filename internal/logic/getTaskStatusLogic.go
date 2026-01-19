package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskStatusLogic {
	return &GetTaskStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取任务状态
func (l *GetTaskStatusLogic) GetTaskStatus(in *spider.GetTaskStatusReq) (*spider.GetTaskStatusResp, error) {
	// todo: add your logic here and delete this line

	return &spider.GetTaskStatusResp{}, nil
}
