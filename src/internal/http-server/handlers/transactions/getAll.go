package transactions

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	trResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/transactions"
)

type IGetter interface {
	GetAll() ([]ents.TransactionTransfer, error)
}

// swagger:route GET /transactions Transaction TransactionsGet
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Schemes: http
//
//
//	 Security:
//		  bearerAuth:
//
//	 Responses:
//	   default: ValidateError
//	   200: TransactionsGetAllResponse
func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.GetAllResponse
		const op = "handlers.transactions.getAll"
		var transactions []ents.TransactionTransfer
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		transactions, err = ctrl.GetAll()
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
