package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test2/internal/logic/user"
	"test2/internal/svc"
)

func UserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
