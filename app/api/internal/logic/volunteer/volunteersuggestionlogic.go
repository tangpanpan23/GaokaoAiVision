package volunteer

import (
	"context"

	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VolunteerSuggestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVolunteerSuggestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VolunteerSuggestionLogic {
	return &VolunteerSuggestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VolunteerSuggestionLogic) VolunteerSuggestion(req *types.VolunteerSuggestionReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
