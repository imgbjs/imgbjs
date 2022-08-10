package logic

import (
	"context"

	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(req *types.UploadImgRequest, downloadURL string) (resp *types.UploadImgResponse, err error) {
	// todo: DB
	return &types.UploadImgResponse{
		DownloadURL: downloadURL,
		ShortURL:    l.svcCtx.Config.ShortNamePrefix + req.ShortName,
	}, nil
}
