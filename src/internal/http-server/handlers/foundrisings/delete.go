package foundrisings

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IDeleter interface {
	Delete(id string) (ents.Foundrising, error)
}

// swagger:route DELETE /foundrisings/{id} Foundrising FoundrisingsDelete
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Schemes: http
//
//
//	 Security:
//		  bearerAuth:
//
//	 Parameters:
//	      + name: id
//	        in: query
//	        required: true
//	        type: integer
//	        format: int64
//
//	 Responses:
//	   default: ValidateError
//	   200: FoundrisingsBaseResponse
//	   401: ValidateError
//	   404: ValidateError
func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		const op = "handlers.foundrisings.delete"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		id := mux.Vars(r)["id"]
		foundrising, err := ctrl.Delete(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundrising: foundrising}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
