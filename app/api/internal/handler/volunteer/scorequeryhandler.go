package volunteer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lighthouse-volunteer/app/api/internal/logic/volunteer"
	"lighthouse-volunteer/app/api/internal/svc"
	"lighthouse-volunteer/app/api/internal/types"
)

func ScoreQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScoreQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := volunteer.NewScoreQueryLogic(r.Context(), svcCtx)
		resp, err := l.ScoreQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
