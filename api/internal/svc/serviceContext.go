package svc

import (
	"github.com/go-redis/redis/v9"
	"github.com/imgbjs/imgbjs/api/internal/config"
	"github.com/imgbjs/imgbjs/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Client
	ImgPoolModel      model.ImgPoolModel
	ShortUrlPoolModel model.ShortUrlPoolModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass}),
		ImgPoolModel:      model.NewImgPoolModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ShortUrlPoolModel: model.NewShortUrlPoolModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
