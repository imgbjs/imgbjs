package handler

import (
	"errors"
	"net/http"

	"github.com/imgbjs/imgbjs/api/internal/logic"
	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RedirectShortURLHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedirectShortURLRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		sup, err := svcCtx.ShortUrlPoolModel.FindOneByShortURL(r.Context(), svcCtx.Config.ShortNamePrefix+req.ShortName)
		if err != nil {
			httpx.Error(w, errors.New("此链接不存在"))
			return
		}
		ip, err := svcCtx.ImgPoolModel.FindOneByIdentity(r.Context(), sup.ImgIdentity)
		if err != nil {
			httpx.Error(w, errors.New("此图片不存在"))
			return
		}
		http.Redirect(w, r, ip.DownloadUrl, http.StatusFound)
		return
		l := logic.NewRedirectShortURLLogic(r.Context(), svcCtx)
		resp, err := l.RedirectShortURL(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
