package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IReplenator interface {
	ReplenishBalance(id string,
		params ents.UserReplenish) (ents.UserTransfer, error)
}

func Replenish(log *slog.Logger, ctrl IReplenator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.BaseResponse
		var user ents.UserTransfer
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		const op = "handlers.user.replenish"
		log.With(
			slog.String("operation", op),
		)
		var params ents.UserReplenish
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		log.Info("Body successfully decoded")
		user, err = ctrl.ReplenishBalance(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = usrResp.BaseResponse{
			User: user,
		}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
