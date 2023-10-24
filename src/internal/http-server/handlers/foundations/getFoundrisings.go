package foundations

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
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
			status := http.StatusBadRequest
			if err == my_errors.ErrorNotExists {
				status = http.StatusNotFound
			}
			if err == my_errors.ErrorConflict {
				status = http.StatusConflict
			}
			log.Error(err.Error())
			resp.JSONRender(w, status, response)
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
		resp.JSONRender(w, http.StatusOK, response)
	}
}
