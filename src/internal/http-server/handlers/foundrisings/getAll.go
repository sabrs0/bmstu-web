package foundrisings

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IGetter interface {
	GetAll() ([]ents.FoundrisingTransfer, error)
}

// swagger:route GET /foundrisings Foundrising FoundrisingsGet
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
//  Schemes: http
//
//
//
//  Responses:
//    default: ValidateError
//    200: FoundrisingsGetAllResponse

func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.GetAllResponse
		var foundrisings []ents.FoundrisingTransfer
		const op = "handlers.foundrisings.getAll"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		foundrisings, err = ctrl.GetAll()
		if err != nil {
			return
		}
		response.Foundrisings = foundrisings
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
