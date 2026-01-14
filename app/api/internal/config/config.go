package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// JWT 配置
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 数据库配置
	DBConfig struct {
		Host         string
		Port         int
		User         string
		Password     string
		Database     string
		MaxIdleConns int `json:",default=10"`
		MaxOpenConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600000"`
	}

	// RPC服务配置
	ScoreRpc  zrpc.RpcClientConf
	SpiderRpc zrpc.RpcClientConf
	AiRpc     zrpc.RpcClientConf
	UserRpc   zrpc.RpcClientConf

	// 缓存配置
	Cache cache.CacheConf
}
