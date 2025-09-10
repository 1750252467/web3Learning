package class

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test2/internal/logic/class"
	"test2/internal/svc"
	"test2/internal/types"
)

func UserClassUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserClassUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewUserClassUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserClassUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
