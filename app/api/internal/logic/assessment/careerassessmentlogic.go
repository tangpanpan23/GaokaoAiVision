package assessment

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CareerAssessmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCareerAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CareerAssessmentLogic {
	return &CareerAssessmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CareerAssessmentLogic) CareerAssessment(req *types.CareerAssessmentReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
