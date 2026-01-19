package assessment

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectPlanQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectPlanQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectPlanQueryLogic {
	return &SubjectPlanQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectPlanQueryLogic) SubjectPlanQuery(req *types.SubjectPlanQueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
