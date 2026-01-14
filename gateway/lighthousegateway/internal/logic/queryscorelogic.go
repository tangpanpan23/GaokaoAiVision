package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"
	"lighthouse-volunteer/services/data/rpc/data"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryScoreLogic {
	return &QueryScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryScoreLogic) QueryScore(req *types.QueryScoreRequest) (resp *types.QueryScoreResponse, err error) {
	// 调用数据查询服务的RPC接口
	rpcReq := &data.QueryScoreReq{
		Year:      int32(req.Year),
		Province:  req.Province,
		College:   req.College,
		Major:     req.Major,
		ScoreType: int32(req.ScoreType),
		Batch:     req.Batch,
		MinScore:  int32(req.MinScore),
		MaxScore:  int32(req.MaxScore),
	}

	rpcResp, err := l.svcCtx.DataRpc.QueryScore(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("Failed to query scores: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp = &types.QueryScoreResponse{
		Code: 200,
		Msg:  "success",
		Total: rpcResp.Total,
	}

	for _, item := range rpcResp.Scores {
		resp.Data = append(resp.Data, types.AdmissionScore{
			Id:          item.Id,
			Year:        int(item.Year),
			Province:    item.Province,
			CollegeName: item.CollegeName,
			MajorName:   item.MajorName,
			ScoreType:   int(item.ScoreType),
			Batch:       item.Batch,
			MinScore:    int(item.MinScore),
			MinRank:     int(item.MinRank),
			DataSource:  item.DataSource,
		})
	}

	return resp, nil
}
