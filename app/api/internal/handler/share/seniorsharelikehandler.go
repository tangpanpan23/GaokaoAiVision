package share

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lighthouse-volunteer/app/api/internal/logic/share"
	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"
)

func SeniorShareLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeniorShareLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := share.NewSeniorShareLikeLogic(r.Context(), svcCtx)
		resp, err := l.SeniorShareLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
