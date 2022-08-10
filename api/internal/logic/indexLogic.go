package logic

import (
	"context"

	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.IndexRequest) (resp *types.IndexResponse, err error) {
	return
}
