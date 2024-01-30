package transactions

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	trResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/transactions"
)

type IByIdGetter interface {
	GetByID(id_ string) (ents.TransactionTransfer, error)
}

// swagger:route GET /transactions/{id} Transaction TransactionsGetById
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
//  Parameters:
//       + name: id
//         in: path
//         required: true
//         type: integer
//         format: int64
//
//  Responses:
//    default: ValidateError
//    200: TransactionsBaseResponse
//    404: ValidateError

func GetByID(log *slog.Logger, ctrl IByIdGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.BaseResponse
		var transaction ents.TransactionTransfer
		const op = "handlers.transactions.getByID"
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
		transaction, err = ctrl.GetByID(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = trResp.BaseResponse{Transaction: transaction}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
