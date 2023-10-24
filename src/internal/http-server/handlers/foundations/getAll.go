package foundations

import (
	"encoding/json"
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IGetter interface {
	GetAll() ([]ents.Foundation, error)
}

func GetAll(logger *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.foundations.getAll"
		logger = slog.With(
			slog.String("operation", op),
		)
		var response fndResp.GetAllResponse
		foundations, err := ctrl.GetAll()
		if err != nil {
			resp.JSONRender(w, http.StatusBadRequest, response)
			return
		}
		response.Foundations = foundations
		data, err := json.Marshal(response)
		resp.JSONRender(w, http.StatusOK, data)
	}
}
