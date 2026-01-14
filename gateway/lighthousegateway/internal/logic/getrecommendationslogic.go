package logic

import (
	"context"

	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"
	"lighthouse-volunteer/services/ai/rpc/ai"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendationsLogic {
	return &GetRecommendationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendationsLogic) GetRecommendations(req *types.AIRecRequest) (resp *types.AIRecResponse, err error) {
	// 调用AI分析服务的RPC接口
	rpcReq := &ai.GetRecommendationsReq{
		Province:  req.Province,
		Subjects:  req.Subjects,
		Score:     int32(req.Score),
		Rank:      int32(req.Rank),
		Interests: req.Interests,
	}

	rpcResp, err := l.svcCtx.AIRpc.GetRecommendations(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("Failed to get AI recommendations: %v", err)
		return nil, err
	}

	// 转换响应格式
	resp = &types.AIRecResponse{
		Code: 200,
		Msg:  "success",
	}

	for _, rec := range rpcResp.Recommendations {
		resp.Data = append(resp.Data, types.Recommendation{
			Level:      rec.Level,
			College:    rec.College,
			Major:      rec.Major,
			MinRank:    int(rec.MinRank),
			Reason:     rec.Reason,
			MatchScore: rec.MatchScore,
		})
	}

	return resp, nil
}
