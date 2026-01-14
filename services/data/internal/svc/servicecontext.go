package svc

import (
	"lighthouse-volunteer/services/data/internal/config"
	"lighthouse-volunteer/services/data/internal/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Redis
	Cache       cache.Cache
	AdmissionModel model.AdmissionScoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.Redis),
		Cache:  cache.New(c.Cache, nil, nil, nil),
		AdmissionModel: model.NewAdmissionScoreModel(c.Database.Host, c.Database.Port,
			c.Database.User, c.Database.Password, c.Database.Database, c.Database.Charset),
	}
}
