package logic

import (
	"context"
	"lighthouse-volunteer/services/data/internal/model"
	"lighthouse-volunteer/services/data/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryScoreLogic {
	return &QueryScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryScoreLogic) QueryScore(req *QueryScoreReq) (*QueryScoreResp, error) {
	// 转换请求格式
	queryReq := &model.QueryScoreReq{
		Year:      int(req.Year),
		Province:  req.Province,
		College:   req.College,
		Major:     req.Major,
		ScoreType: int(req.ScoreType),
		Batch:     req.Batch,
		MinScore:  int(req.MinScore),
		MaxScore:  int(req.MaxScore),
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
	}

	// 查询数据
	scores, err := l.svcCtx.AdmissionModel.FindByCondition(queryReq)
	if err != nil {
		l.Logger.Errorf("Failed to query scores: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp := &QueryScoreResp{
		Code: 200,
		Msg:  "success",
	}

	for _, score := range scores {
		resp.Scores = append(resp.Scores, &AdmissionScore{
			Id:          score.Id,
			Year:        int32(score.Year),
			Province:    score.Province,
			CollegeName: score.CollegeName,
			MajorName:   score.MajorName,
			ScoreType:   int32(score.ScoreType),
			Batch:       score.Batch,
			MinScore:    int32(score.MinScore),
			MinRank:     int32(score.MinRank),
			DataSource:  score.DataSource,
		})
	}

	// 获取总数（这里简化处理，实际应该有单独的计数查询）
	resp.Total = int64(len(scores))

	return resp, nil
}
