package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"task4/internal/logic/comment"
	"task4/internal/svc"
	"task4/internal/types"
)

// 获取评论列表
func GetCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewGetCommentsLogic(r.Context(), svcCtx)
		resp, err := l.GetComments(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
