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

// swagger:operation GET /foundations/{id} FoundationsGetById
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - text/plain
//
// parameters:
//   - name: id
//     in: path
//     required: true
//     schema:
//     type: integer
//     format: int32
//
// responses:
//
// '200':
//
//	description: Success
//	schema:
//	 "$ref": "#/definitions/Foundation"
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
