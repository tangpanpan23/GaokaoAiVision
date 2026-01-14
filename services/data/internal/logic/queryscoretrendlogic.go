package logic

import (
	"context"
	"lighthouse-volunteer/services/data/internal/model"
	"lighthouse-volunteer/services/data/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryScoreTrendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryScoreTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryScoreTrendLogic {
	return &QueryScoreTrendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryScoreTrendLogic) QueryScoreTrend(req *QueryScoreTrendReq) (*QueryScoreTrendResp, error) {
	// 转换请求格式
	queryReq := &model.QueryScoreTrendReq{
		College:   req.College,
		Major:     req.Major,
		Province:  req.Province,
		StartYear: int(req.StartYear),
		EndYear:   int(req.EndYear),
	}

	// 查询数据
	trends, err := l.svcCtx.AdmissionModel.FindTrends(queryReq)
	if err != nil {
		l.Logger.Errorf("Failed to query score trends: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp := &QueryScoreTrendResp{
		Code: 200,
		Msg:  "success",
	}

	for _, trend := range trends {
		resp.Trends = append(resp.Trends, &ScoreTrend{
			Year:     int32(trend.Year),
			College:  trend.College,
			Major:    trend.Major,
			MinScore: int32(trend.MinScore),
			MinRank:  int32(trend.MinRank),
			AvgScore: trend.AvgScore,
			Batch:    trend.Batch,
		})
	}

	return resp, nil
}
