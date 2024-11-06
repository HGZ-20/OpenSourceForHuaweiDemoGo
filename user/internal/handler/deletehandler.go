package handler

import (
	"net/http"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/logic"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
