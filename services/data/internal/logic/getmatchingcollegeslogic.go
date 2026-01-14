package logic

import (
	"context"
	"lighthouse-volunteer/services/data/internal/model"
	"lighthouse-volunteer/services/data/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMatchingCollegesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMatchingCollegesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMatchingCollegesLogic {
	return &GetMatchingCollegesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMatchingCollegesLogic) GetMatchingColleges(req *GetMatchingCollegesReq) (*GetMatchingCollegesResp, error) {
	// 转换请求格式
	queryReq := &model.GetMatchingCollegesReq{
		Score:     int(req.Score),
		Rank:      int(req.Rank),
		Province:  req.Province,
		Subjects:  req.Subjects,
		ScoreType: int(req.ScoreType),
		Batch:     req.Batch,
		Limit:     int(req.Limit),
	}

	// 查询数据
	colleges, err := l.svcCtx.AdmissionModel.FindMatchingColleges(queryReq)
	if err != nil {
		l.Logger.Errorf("Failed to get matching colleges: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp := &GetMatchingCollegesResp{
		Code: 200,
		Msg:  "success",
	}

	for _, college := range colleges {
		resp.Colleges = append(resp.Colleges, &MatchingCollege{
			CollegeName: college.CollegeName,
			MajorName:   college.MajorName,
			MinScore:    int32(college.MinScore),
			MinRank:     int32(college.MinRank),
			Batch:       college.Batch,
			Year:        int32(college.Year),
			Probability: college.Probability,
		})
	}

	return resp, nil
}
