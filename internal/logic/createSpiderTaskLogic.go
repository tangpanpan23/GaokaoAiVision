package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpiderTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSpiderTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpiderTaskLogic {
	return &CreateSpiderTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建爬虫任务
func (l *CreateSpiderTaskLogic) CreateSpiderTask(in *spider.CreateSpiderTaskReq) (*spider.CreateSpiderTaskResp, error) {
	// todo: add your logic here and delete this line

	return &spider.CreateSpiderTaskResp{}, nil
}
