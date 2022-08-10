package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/imgbjs/imgbjs/api/internal/logic"
	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"
	"github.com/imgbjs/imgbjs/model"
	"github.com/imgbjs/imgbjs/util"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		var err error
		n := r.FormValue("shortname")
		if len(n) > 0 {
			if sup, err := svcCtx.ShortUrlPoolModel.FindOneByShortURL(r.Context(), svcCtx.Config.ShortNamePrefix+n); err == nil && sup != nil {
				httpx.Error(w, errors.New("此短链已被使用"))
				return
			}
		}
		if err != nil && err != model.ErrNotFound {
			httpx.Error(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		req.Hash = hash
		req.Ext = path.Ext(fileHeader.Filename)
		req.Name = fileHeader.Filename
		req.Size = fileHeader.Size
		identity := util.NewUuid() // d46b416d-3428-44a8-84d5-acaaf9798ae0
		shortName := identity[:8]
		req.ShortName = shortName[:8]
		if len(n) > 0 {
			req.ShortName = n
		}
		shortURL := svcCtx.Config.ShortNamePrefix + req.ShortName
		if ip, err := svcCtx.ImgPoolModel.FindOneByHash(r.Context(), hash); err == nil && ip != nil {
			if _, err = svcCtx.ShortUrlPoolModel.InsertUpdateCacheShortURL(r.Context(), &model.ShortUrlPool{
				ImgIdentity: ip.Identity,
				ShortUrl:    shortURL,
				DownloadUrl: ip.DownloadUrl,
			}); err != nil {
				httpx.Error(w, err)
				return
			}
			httpx.OkJson(w, &types.UploadImgResponse{
				DownloadURL: ip.DownloadUrl,
				ShortURL:    shortURL,
			})
			return
		} else if err != nil && err != model.ErrNotFound {
			httpx.Error(w, err)
			return
		}

		downloadURL, err := util.UploadGithub(shortName, identity+req.Ext, b)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if _, err = svcCtx.ImgPoolModel.Insert(r.Context(), &model.ImgPool{
			Identity:    identity,
			Hash:        req.Hash,
			Ext:         req.Ext,
			Size:        req.Size,
			DownloadUrl: downloadURL,
		}); err != nil {
			httpx.Error(w, err)
			return
		}
		if _, err = svcCtx.ShortUrlPoolModel.InsertUpdateCacheShortURL(r.Context(), &model.ShortUrlPool{
			ImgIdentity: identity,
			ShortUrl:    shortURL,
			DownloadUrl: downloadURL,
		}); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(&req, downloadURL)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
