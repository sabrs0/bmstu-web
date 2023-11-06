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

type IDonator interface {
	DonateToFoundrising(foundationID string,
		params ents.FoundationDonate) (ents.Transaction, error)
}

// swagger:route POST /foundations/{id}/donate Foundation FoundationsDonate
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
//  Schemes: http
//
//
//  Security:
//	  bearerAuth:
//
//  Responses:
//    default: ValidateError
//    200: FoundationsDonateResponse
//    400: ValidateError
//    401: ValidateError
//    404: ValidateError
//    409: ValidateError

func Donate(log *slog.Logger, ctrl IDonator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.DonateResponse
		var transaction ents.Transaction
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

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
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
