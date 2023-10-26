package foundrisings

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IByIdGetter interface {
	GetByID(id_ string) (ents.Foundrising, error)
}

func GetByID(log *slog.Logger, ctrl IByIdGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		const op = "handlers.foundrisings.getByID"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				resp.ErrWrapper(log, w, &response, err)
			}

		}()
		id := mux.Vars(r)["id"]
		foundrising, err := ctrl.GetByID(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundrising: foundrising}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
