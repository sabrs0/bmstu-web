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

// swagger:operation DELETE /foundations/{id} FoundationsDelete
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
//	'200':
//	  description: Success
//	  schema:
//	    "$ref": "#/definitions/Foundation"
func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
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
