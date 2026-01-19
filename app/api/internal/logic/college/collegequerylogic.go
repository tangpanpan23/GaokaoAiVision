package college

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollegeQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollegeQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollegeQueryLogic {
	return &CollegeQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollegeQueryLogic) CollegeQuery(req *types.CollegeQueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
