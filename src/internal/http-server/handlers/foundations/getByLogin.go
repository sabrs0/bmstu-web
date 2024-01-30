package foundations

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IByLoginGetter interface {
	GetExtByLogin(login string) (ents.FoundationTransferExtended, error)
}

func GetExtByLogin(log *slog.Logger, ctrl IByLoginGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.ExtResponse
		var foundation ents.FoundationTransferExtended
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		const op = "handlers.foundation.getByLogin"
		log = log.With(
			slog.String("operation", op),
		)
		login := mux.Vars(r)["login"]
		foundation, err = ctrl.GetExtByLogin(login)
		if err != nil {
			return
		}
		response = fndResp.ExtResponse{Foundation: foundation}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)

	}
}
