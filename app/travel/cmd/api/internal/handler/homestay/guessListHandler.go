package homestay

import (
	"net/http"

	"automatix/app/travel/cmd/api/internal/logic/homestay"
	"automatix/app/travel/cmd/api/internal/svc"
	"automatix/app/travel/cmd/api/internal/types"
	"automatix/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GuessListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestay.NewGuessListLogic(r.Context(), ctx)
		resp, err := l.GuessList(req)
		result.HttpResult(r, w, resp, err)
	}
}
