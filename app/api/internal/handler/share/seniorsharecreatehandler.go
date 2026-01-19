package share

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lighthouse-volunteer/app/api/internal/logic/share"
	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"
)

func SeniorShareCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeniorShareCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := share.NewSeniorShareCreateLogic(r.Context(), svcCtx)
		resp, err := l.SeniorShareCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
