package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `json:",default=localhost"`
	Port     int    `json:",default=3306"`
	User     string `json:",default=root"`
	Password string `json:",default=password"`
	Database string `json:",default=lighthouse_volunteer"`
	Charset  string `json:",default=utf8mb4"`
}

// JwtConfig JWT配置
type JwtConfig struct {
	Secret  string `json:",default=your-jwt-secret-key"`
	Expire  int64  `json:",default=86400"` // 秒
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `json:",default=info"`
	Encoding string `json:",default=json"`
}

type Config struct {
	rest.RestConf

	// 数据库配置
	Database DatabaseConfig `json:",optional"`

	// Redis配置
	Redis redis.RedisConf `json:",optional"`

	// RPC服务配置
	RpcClients struct {
		DataRpc   zrpc.RpcClientConf `json:",optional"`
		AIRpc     zrpc.RpcClientConf `json:",optional"`
		UserRpc   zrpc.RpcClientConf `json:",optional"`
		SpiderRpc zrpc.RpcClientConf `json:",optional"`
	} `json:",optional"`

	// JWT配置
	Jwt JwtConfig `json:",optional"`

	// 日志配置
	Log LogConfig `json:",optional"`
}
