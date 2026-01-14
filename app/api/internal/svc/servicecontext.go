package svc

import (
	"lighthouse-volunteer/app/api/internal/config"
	"lighthouse-volunteer/app/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	CORS        rest.Middleware
	JwtAuth     rest.Middleware
	RateLimit   rest.Middleware

	// RPC 客户端
	ScoreRpc  score.Score
	SpiderRpc spider.Spider
	AiRpc     ai.Ai
	UserRpc   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		CORS:      middleware.NewCORSMiddleware().Handle,
		JwtAuth:   middleware.NewJwtAuthMiddleware(c.JwtAuth.AccessSecret).Handle,
		RateLimit: middleware.NewRateLimitMiddleware().Handle,

		// 初始化RPC客户端
		ScoreRpc:  score.NewScore(zrpc.MustNewClient(c.ScoreRpc)),
		SpiderRpc: spider.NewSpider(zrpc.MustNewClient(c.SpiderRpc)),
		AiRpc:     ai.NewAi(zrpc.MustNewClient(c.AiRpc)),
		UserRpc:   user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
