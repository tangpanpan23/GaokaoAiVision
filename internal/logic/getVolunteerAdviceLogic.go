package logic

import (
	"context"

	"lighthouse-volunteer/internal/svc"
	"lighthouse-volunteer/pb/lighthouse-volunteer/app/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVolunteerAdviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVolunteerAdviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVolunteerAdviceLogic {
	return &GetVolunteerAdviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AI志愿咨询
func (l *GetVolunteerAdviceLogic) GetVolunteerAdvice(in *ai.GetVolunteerAdviceReq) (*ai.GetVolunteerAdviceResp, error) {
	// todo: add your logic here and delete this line

	return &ai.GetVolunteerAdviceResp{}, nil
}
