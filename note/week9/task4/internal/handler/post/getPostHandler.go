package post

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"task4/internal/logic/post"
	"task4/internal/svc"
	"task4/internal/types"
)

// 获取文章详情
func GetPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := post.NewGetPostLogic(r.Context(), svcCtx)
		resp, err := l.GetPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
