package foundations

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IFoundrisingsGetter interface {
	GetAllFoundrisings(id string) ([]ents.Foundrising, error)
}

func GetFoundrisings(log *slog.Logger, ctrl IFoundrisingsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.FoundrisingsResponse
		var foundrisings []ents.Foundrising
		defer func() {
			resp.ErrWrapper(log, w, response, err)
		}()
		const op = "handlers.foundation.getFoundrisings"
		log = slog.With(
			slog.String("operation", op),
		)
		id := mux.Vars(r)["id"]
		foundrisings, err = ctrl.GetAllFoundrisings(id)
		if err != nil {
			return
		}
		response = fndResp.FoundrisingsResponse{Foundrisings: foundrisings}
		log.Info("Successs")
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
