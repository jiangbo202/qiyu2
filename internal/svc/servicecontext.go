package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_t2/internal/config"
	"go_zero_t2/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Author model.AuthorModel

	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Author: model.NewAuthorModel(sqlx.NewMysql(c.DB.DataSource)),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr: c.Redis.Host,
			Password: c.Redis.Pass,
		}),
	}
}
