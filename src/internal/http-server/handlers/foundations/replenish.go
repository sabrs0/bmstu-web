package foundations

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IReplenator interface {
	ReplenishBalance(id string,
		params ents.FoundationReplenish) (ents.FoundationTransfer, error)
}

func Replenish(log *slog.Logger, ctrl IReplenator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		var foundation ents.FoundationTransfer
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		const op = "handlers.foundation.replenish"
		log.With(
			slog.String("operation", op),
		)
		var params ents.FoundationReplenish
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		log.Info("Body successfully decoded")
		foundation, err = ctrl.ReplenishBalance(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{
			Foundation: foundation,
		}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
