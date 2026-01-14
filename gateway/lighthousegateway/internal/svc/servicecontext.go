package svc

import (
	"lighthouse-volunteer/gateway/lighthousegateway/internal/config"
	"lighthouse-volunteer/services/data/rpc/data"
	"lighthouse-volunteer/services/ai/rpc/ai"
	"lighthouse-volunteer/services/user/rpc/user"
	"lighthouse-volunteer/services/spider/rpc/spider"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// RPC客户端
	DataRpc   data.DataClient
	AIRpc     ai.AiClient
	UserRpc   user.UserClient
	SpiderRpc spider.SpiderClient

	// Redis客户端
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 初始化RPC客户端
		DataRpc:   data.NewDataClient(zrpc.MustNewClient(c.RpcClients.DataRpc)),
		AIRpc:     ai.NewAiClient(zrpc.MustNewClient(c.RpcClients.AIRpc)),
		UserRpc:   user.NewUserClient(zrpc.MustNewClient(c.RpcClients.UserRpc)),
		SpiderRpc: spider.NewSpiderClient(zrpc.MustNewClient(c.RpcClients.SpiderRpc)),

		// 初始化Redis客户端
		Redis: redis.MustNewRedis(c.Redis),
	}
}
