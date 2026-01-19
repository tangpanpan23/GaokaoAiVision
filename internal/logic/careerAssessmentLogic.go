package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CareerAssessmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCareerAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CareerAssessmentLogic {
	return &CareerAssessmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 职业测评
func (l *CareerAssessmentLogic) CareerAssessment(in *user.CareerAssessmentReq) (*user.CareerAssessmentResp, error) {
	// todo: add your logic here and delete this line

	return &user.CareerAssessmentResp{}, nil
}
