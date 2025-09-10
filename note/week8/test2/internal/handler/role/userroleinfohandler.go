package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test2/internal/logic/role"
	"test2/internal/svc"
	"test2/internal/types"
)

func UserRoleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewUserRoleInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserRoleInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
