package transactions

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	trResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/transactions"
)

type IGetter interface {
	GetAll() ([]ents.Transaction, error)
}

// swagger:operation GET /transactions TransactionsGet
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - text/plain
// responses:
//
//		'200':
//		schema:
//			type: array
//			items:
//	 		$ref": "#/definitions/Transaction"
//		  description: Success
func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.GetAllResponse
		const op = "handlers.transactions.getAll"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		transactions, err := ctrl.GetAll()
		if err != nil {
			return
		}
		response = trResp.GetAllResponse{
			Transactions: transactions,
		}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
