package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/logic"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/svc"
	"lighthouse-volunteer/gateway/lighthousegateway/internal/types"
)

func AIChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AIChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAIChatLogic(r.Context(), svcCtx)
		resp, err := l.AIChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
