package handler

import (
	"net/http"
	"path"
	"text/template"

	"github.com/imgbjs/imgbjs/api/internal/logic"
	"github.com/imgbjs/imgbjs/api/internal/svc"
	"github.com/imgbjs/imgbjs/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Game2048Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(path.Join("view", "2048.html"))
		t.Execute(w, nil)
		return
		var req types.IndexRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGame2048Logic(r.Context(), svcCtx)
		resp, err := l.Game2048(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
