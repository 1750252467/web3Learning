package blog

import (
	"net/http"

	"blog/api/api/internal/logic/blog"
	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取文章列表
func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewGetPostsLogic(r.Context(), svcCtx)
		resp, err := l.GetPosts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
