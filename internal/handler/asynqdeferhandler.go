package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_t2/internal/logic"
	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"
)

func AsynqDeferHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AsynqDeferTask1Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAsynqDeferLogic(r.Context(), svcCtx)
		resp, err := l.AsynqDefer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
