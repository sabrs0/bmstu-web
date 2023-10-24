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

type IDeleter interface {
	Delete(id string) (ents.Foundation, error)
}

func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
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
		resp.JSONRender(w, http.StatusOK, response)

	}
}
