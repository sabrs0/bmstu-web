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

// swagger:operation DELETE /foundrisings/{id} Foundrising FoundrisingsDelete
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
//  '200':
//   description: Success
//   schema:
//    "$ref": "#/definitions/Foundrising"
//  '401':
//   description: Unauthorized
//  '404':
//   description: Not Found
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
