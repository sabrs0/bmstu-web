package foundations

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IUpdater interface {
	Update(id string, params ents.FoundationAdd) (ents.Foundation, error)
}

func Update(log *slog.Logger, ctrl IUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		defer func() {
			if err != nil {
				status := http.StatusBadRequest
				if errors.Is(err, my_errors.ErrorNotExists) {
					status = http.StatusNotFound
				}
				if errors.Is(err, my_errors.ErrorConflict) {
					status = http.StatusConflict
				}
				log.Error(err.Error())
				resp.JSONRender(w, status, response)

			}
		}()
		const op = "handlers.foundations.update"
		log = log.With(
			slog.String("operation", op),
		)
		vars := mux.Vars(r)
		id := vars["id"]
		var params ents.FoundationAdd
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		log.Info("Request body decoded")
		var foundation ents.Foundation
		foundation, err = ctrl.Update(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundation: foundation}
		resp.JSONRender(w, http.StatusOK, response)
	}
}
