// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shortUrlPoolFieldNames          = builder.RawFieldNames(&ShortUrlPool{})
	shortUrlPoolRows                = strings.Join(shortUrlPoolFieldNames, ",")
	shortUrlPoolRowsExpectAutoSet   = strings.Join(stringx.Remove(shortUrlPoolFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	shortUrlPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(shortUrlPoolFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheImgbjsShortUrlPoolIdPrefix = "cache:imgbjs:shortUrlPool:id:"
	cacheImgbjsShortUrlPoolImgIdentityPrefix = "cache:imgbjs:shortUrlPool:imgIdentity:"
	cacheImgbjsShortUrlPoolShortURLPrefix = "cache:imgbjs:shortUrlPool:shortURL:"
)

type (
	shortUrlPoolModel interface {
		Insert(ctx context.Context, data *ShortUrlPool) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ShortUrlPool, error)
		Update(ctx context.Context, data *ShortUrlPool) error
		Delete(ctx context.Context, id int64) error
		FindOneByImgIdentity(ctx context.Context, imgIdentity string) (*ShortUrlPool, error)
		FindOneByShortURL(ctx context.Context, shortURL string) (*ShortUrlPool, error)
		InsertUpdateCacheShortURL(ctx context.Context, data *ShortUrlPool) (sql.Result, error)
	}

	defaultShortUrlPoolModel struct {
		sqlc.CachedConn
		table string
	}

	ShortUrlPool struct {
		Id          int64     `db:"id"`
		ImgIdentity string    `db:"img_identity"`
		ShortUrl    string    `db:"short_url"`
		DownloadUrl string    `db:"download_url"`
		CreateTime  time.Time `db:"create_time"`
	}
)

func newShortUrlPoolModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultShortUrlPoolModel {
	return &defaultShortUrlPoolModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`short_url_pool`",
	}
}

func (m *defaultShortUrlPoolModel) Delete(ctx context.Context, id int64) error {
	imgbjsShortUrlPoolIdKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, imgbjsShortUrlPoolIdKey)
	return err
}

func (m *defaultShortUrlPoolModel) FindOne(ctx context.Context, id int64) (*ShortUrlPool, error) {
	imgbjsShortUrlPoolIdKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolIdPrefix, id)
	var resp ShortUrlPool
	err := m.QueryRowCtx(ctx, &resp, imgbjsShortUrlPoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shortUrlPoolRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlPoolModel) Insert(ctx context.Context, data *ShortUrlPool) (sql.Result, error) {
	imgbjsShortUrlPoolIdKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, shortUrlPoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.ImgIdentity, data.ShortUrl, data.DownloadUrl)
	}, imgbjsShortUrlPoolIdKey)
	return ret, err
}

func (m *defaultShortUrlPoolModel) Update(ctx context.Context, data *ShortUrlPool) error {
	imgbjsShortUrlPoolIdKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shortUrlPoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ImgIdentity, data.ShortUrl, data.DownloadUrl, data.Id)
	}, imgbjsShortUrlPoolIdKey)
	return err
}

func (m *defaultShortUrlPoolModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolIdPrefix, primary)
}

func (m *defaultShortUrlPoolModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shortUrlPoolRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShortUrlPoolModel) tableName() string {
	return m.table
}

func (m *defaultShortUrlPoolModel) FindOneByImgIdentity(ctx context.Context, imgIdentity string) (*ShortUrlPool, error) {
	imgbjsShortUrlPoolImgIdentityKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolImgIdentityPrefix, imgIdentity)
	var resp ShortUrlPool
	err := m.QueryRowCtx(ctx, &resp, imgbjsShortUrlPoolImgIdentityKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `img_identity` = ? limit 1", shortUrlPoolRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, imgIdentity)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlPoolModel) FindOneByShortURL(ctx context.Context, shortURL string) (*ShortUrlPool, error) {
	imgbjsShortUrlPoolShortURLKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolShortURLPrefix, shortURL)
	var resp ShortUrlPool
	err := m.QueryRowCtx(ctx, &resp, imgbjsShortUrlPoolShortURLKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `short_url` = ? limit 1", shortUrlPoolRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, shortURL)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlPoolModel) InsertUpdateCacheShortURL(ctx context.Context, data *ShortUrlPool) (sql.Result, error) {
	cacheImgbjsShortUrlPoolShortURLPrefixKey := fmt.Sprintf("%s%v", cacheImgbjsShortUrlPoolShortURLPrefix, data.ShortUrl)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, shortUrlPoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.ImgIdentity, data.ShortUrl, data.DownloadUrl)
	}, cacheImgbjsShortUrlPoolShortURLPrefixKey)
	return ret, err
}