package foundations

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IByIdGetter interface {
	GetByID(id_ string) (ents.Foundation, error)
}

func GetByID(log *slog.Logger, ctrl IByIdGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		var foundation ents.Foundation
		defer func() {
			status := http.StatusBadRequest
			if errors.Is(err, my_errors.ErrorNotExists) {
				status = http.StatusNotFound
			}
			if errors.Is(err, my_errors.ErrorConflict) {
				status = http.StatusConflict
			}
			log.Error(err.Error())
			resp.JSONRender(w, status, response)
		}()
		const op = "handlers.foundation.getById"
		log = log.With(
			slog.String("operation", op),
		)
		id := mux.Vars(r)["id"]
		foundation, err = ctrl.GetByID(id)
		if err != nil {
			return
		}
		response = fndResp.BaseResponse{Foundation: foundation}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, response)

	}
}
