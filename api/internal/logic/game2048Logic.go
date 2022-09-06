package logic

import (
	"context"

	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Game2048Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGame2048Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Game2048Logic {
	return &Game2048Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Game2048Logic) Game2048(req *types.IndexRequest) (resp *types.IndexResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
