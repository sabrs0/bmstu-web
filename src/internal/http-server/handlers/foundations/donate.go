package foundations

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IDonator interface {
	DonateToFoundrising(foundationID string,
		params ents.FoundationDonate) (ents.Transaction, error)
}

func Donate(log *slog.Logger, ctrl IDonator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.DonateResponse
		var transaction ents.Transaction
		defer func() {
			status := http.StatusBadRequest
			if err == my_errors.ErrorNotExists {
				status = http.StatusNotFound
			}
			if err == my_errors.ErrorConflict {
				status = http.StatusConflict
			}
			log.Error(err.Error())
			resp.JSONRender(w, status, response)
		}()
		const op = "handlers.foundation.donate"
		var params ents.FoundationDonate
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		log.Info("Body successfully decoded")
		transaction, err = ctrl.DonateToFoundrising(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.DonateResponse{Transaction: transaction}
		resp.JSONRender(w, http.StatusOK, response)
	}
}
