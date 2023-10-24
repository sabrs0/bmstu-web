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

func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response trResp.BaseResponse
		const op = "handlers.transactions.delete"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			resp.ErrWrapper(log, w, response, err)
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
