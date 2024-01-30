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

type IDonator interface {
	Donate(userID string, params ents.UserDonate) (ents.TransactionTransfer, error)
}

// swagger:route POST /users/{id}/donate User UsersDonate
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
//    200: UsersDonateResponse
//    400: ValidateError
//    401: ValidateError
//    404: ValidateError
//    409: ValidateError

func Donate(log *slog.Logger, ctrl IDonator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.DonateResponse
		var transaction ents.TransactionTransfer
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		const op = "handlers.users.donate"
		var params ents.UserDonate
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		log.Info("Body successfully decoded")
		transaction, err = ctrl.Donate(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = usrResp.DonateResponse{Transaction: transaction}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
