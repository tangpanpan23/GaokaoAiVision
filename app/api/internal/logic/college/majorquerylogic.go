package college

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MajorQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMajorQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MajorQueryLogic {
	return &MajorQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MajorQueryLogic) MajorQuery(req *types.MajorQueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
