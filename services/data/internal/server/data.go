package server

import (
	"context"
	"lighthouse-volunteer/services/data/internal/logic"
	"lighthouse-volunteer/services/data/internal/svc"
)

type DataServer struct {
	svcCtx *svc.ServiceContext
}

func NewDataServer(svcCtx *svc.ServiceContext) *DataServer {
	return &DataServer{
		svcCtx: svcCtx,
	}
}

func (s *DataServer) QueryScore(ctx context.Context, req *QueryScoreReq) (*QueryScoreResp, error) {
	l := logic.NewQueryScoreLogic(ctx, s.svcCtx)
	return l.QueryScore(req)
}

func (s *DataServer) QueryScoreTrend(ctx context.Context, req *QueryScoreTrendReq) (*QueryScoreTrendResp, error) {
	l := logic.NewQueryScoreTrendLogic(ctx, s.svcCtx)
	return l.QueryScoreTrend(req)
}

func (s *DataServer) GetMatchingColleges(ctx context.Context, req *GetMatchingCollegesReq) (*GetMatchingCollegesResp, error) {
	l := logic.NewGetMatchingCollegesLogic(ctx, s.svcCtx)
	return l.GetMatchingColleges(req)
}
