package foundations

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IDeleter interface {
	Delete(id string) (ents.Foundation, error)
}

func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		var foundation ents.Foundation
		defer func() {
			if err != nil {
				resp.ErrWrapper(log, w, &response, err)
			}

		}()
		const op = "handlers.foundations.delete"
		log.With(
			slog.String("operation", op),
		)
		id := mux.Vars(r)["id"]
		foundation, err = ctrl.Delete(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundation: foundation}
		resp.JSONRender(w, http.StatusOK, &response)

	}
}
