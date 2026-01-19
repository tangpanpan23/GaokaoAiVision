package svc

import (
	"lighthouse-volunteer/app/api/internal/config"
	"lighthouse-volunteer/app/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	CORS        rest.Middleware
	JwtAuth     rest.Middleware
	RateLimit   rest.Middleware

	// RPC 客户端（暂时注释，待RPC服务实现后再启用）
	// ScoreRpc  score.ScoreClient
	// SpiderRpc spider.SpiderClient
	// AiRpc     ai.AiClient
	// UserRpc   user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		CORS:      middleware.NewCORSMiddleware().Handle,
		JwtAuth:   middleware.NewJwtAuthMiddleware(c.JwtAuth.AccessSecret).Handle,
		RateLimit: middleware.NewRateLimitMiddleware().Handle,

		// 初始化RPC客户端（暂时注释，待RPC服务实现后再启用）
		// ScoreRpc:  score.NewScoreClient(zrpc.MustNewClient(c.ScoreRpc).Conn()),
		// SpiderRpc: spider.NewSpiderClient(zrpc.MustNewClient(c.SpiderRpc).Conn()),
		// AiRpc:     ai.NewAiClient(zrpc.MustNewClient(c.AiRpc).Conn()),
		// UserRpc:   user.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}

