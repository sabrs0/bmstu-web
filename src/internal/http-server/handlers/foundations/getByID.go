package foundations

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IByIdGetter interface {
	GetByID(id_ string) (ents.Foundation, error)
}

// swagger:route GET /foundations/{id} Foundation FoundationsGetById
//
//		 Consumes:
//		 - application/json
//
//		 Produces:
//		 - application/json
//
//		 Schemes: http
//
//
//
//	 Parameters:
//	      + name: id
//	        in: query
//	        required: true
//	        type: integer
//	        format: int64
//
//		 Responses:
//		   default: ValidateError
//		   200: FoundationsBaseResponse
//		   404: ValidateError
func GetByID(log *slog.Logger, ctrl IByIdGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		var foundation ents.Foundation
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

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
		resp.JSONRender(w, http.StatusOK, &response)

	}
}
