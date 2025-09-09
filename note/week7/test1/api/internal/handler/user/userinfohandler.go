package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/core/logx"
	"note/week7/test1/api/internal/logic/user"
	"note/week7/test1/api/internal/svc"
	"note/week7/test1/api/internal/types"
)

// 获取用户信息
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		// 添加调试日志
		logx.Info("Received UserInfo request", logx.Field("url", r.URL.String()))
		logx.Info("All query parameters", logx.Field("params", r.URL.Query()))
		userIdStr := r.URL.Query().Get("userId")
		logx.Info("Raw userId parameter", logx.Field("value", userIdStr), logx.Field("length", len(userIdStr)))
		
		// 对于GET请求，手动解析userId参数
		if r.Method == "GET" {
			if userIdStr == "" {
				logx.Error("userId parameter is empty")
				httpx.ErrorCtx(r.Context(), w, errors.New("userId参数不能为空"))
				return
			}
			
			// 手动转换字符串为int64
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				logx.Error("Failed to parse userId to int64", logx.Field("error", err.Error()))
				httpx.ErrorCtx(r.Context(), w, errors.New("userId格式错误"))
				return
			}
			
			req.UserId = userId
			logx.Info("Manually parsed userId successfully", logx.Field("userId", req.UserId))
		} else {
			// 对于POST请求，使用标准解析
			if err := httpx.Parse(r, &req); err != nil {
				logx.Error("Failed to parse request", logx.Field("error", err.Error()))
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			logx.Info("Parsed request successfully", logx.Field("userId", req.UserId))
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
