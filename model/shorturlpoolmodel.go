package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShortUrlPoolModel = (*customShortUrlPoolModel)(nil)

type (
	// ShortUrlPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortUrlPoolModel.
	ShortUrlPoolModel interface {
		shortUrlPoolModel
	}

	customShortUrlPoolModel struct {
		*defaultShortUrlPoolModel
	}
)

// NewShortUrlPoolModel returns a model for the database table.
func NewShortUrlPoolModel(conn sqlx.SqlConn, c cache.CacheConf) ShortUrlPoolModel {
	return &customShortUrlPoolModel{
		defaultShortUrlPoolModel: newShortUrlPoolModel(conn, c),
	}
}
