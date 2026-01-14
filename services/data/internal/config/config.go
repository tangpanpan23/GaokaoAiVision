package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

// ElasticsearchConfig Elasticsearch配置
type ElasticsearchConfig struct {
	Addresses []string `json:",optional"`
	Username  string   `json:",optional"`
	Password  string   `json:",optional"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `json:",default=info"`
	Encoding string `json:",default=json"`
}

type Config struct {
	zrpc.RpcServerConf

	// 数据库配置
	Database DatabaseConfig `json:",optional"`

	// Redis配置
	Redis redis.RedisConf `json:",optional"`

	// Elasticsearch配置
	Elasticsearch ElasticsearchConfig `json:",optional"`

	// 缓存配置
	Cache cache.CacheConf `json:",optional"`

	// 日志配置
	Log LogConfig `json:",optional"`
}
