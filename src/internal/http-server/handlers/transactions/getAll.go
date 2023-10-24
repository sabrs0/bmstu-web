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

func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.GetAllResponse
		const op = "handlers.transactions.getAll"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			resp.ErrWrapper(log, w, response, err)
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
