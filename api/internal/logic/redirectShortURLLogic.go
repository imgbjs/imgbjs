package logic

import (
	"context"

	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectShortURLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectShortURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectShortURLLogic {
	return &RedirectShortURLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectShortURLLogic) RedirectShortURL(req *types.RedirectShortURLRequest) (resp *types.RedirectShortURLResponse, err error) {
	return
}
