package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ImgPoolModel = (*customImgPoolModel)(nil)

type (
	// ImgPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImgPoolModel.
	ImgPoolModel interface {
		imgPoolModel
	}

	customImgPoolModel struct {
		*defaultImgPoolModel
	}
)

// NewImgPoolModel returns a model for the database table.
func NewImgPoolModel(conn sqlx.SqlConn, c cache.CacheConf) ImgPoolModel {
	return &customImgPoolModel{
		defaultImgPoolModel: newImgPoolModel(conn, c),
	}
}
