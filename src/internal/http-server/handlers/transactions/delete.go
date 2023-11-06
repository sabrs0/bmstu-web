package transactions

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	trResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/transactions"
)

type IDeleter interface {
	Delete(id string) (ents.Transaction, error)
}

// swagger:operation DELETE /transactions/{id} TransactionsDelete
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - text/plain
//
// parameters:
//   - name: id
//     in: path
//     required: true
//     schema:
//     type: integer
//     format: int32
//
// responses:
//
//	'200':
//	  description: Success
//	  schema:
//	    "$ref": "#/definitions/Transaction"
func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.BaseResponse
		const op = "handlers.transactions.delete"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		id := mux.Vars(r)["id"]
		transaction, err := ctrl.Delete(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = trResp.BaseResponse{Transaction: transaction}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
