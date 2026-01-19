package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectPlanLogic {
	return &GetSubjectPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取选科规划
func (l *GetSubjectPlanLogic) GetSubjectPlan(in *user.GetSubjectPlanReq) (*user.GetSubjectPlanResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetSubjectPlanResp{}, nil
}
