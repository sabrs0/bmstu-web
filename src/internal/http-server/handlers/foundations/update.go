package foundations

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IUpdater interface {
	Update(id string, params ents.FoundationAdd) (ents.Foundation, error)
}

// swagger:operation PUT /foundations/{id} Foundation FoundationsUpdate
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
// - name: id
//   in: path
//   required: true
//   schema:
//   type: integer
//   format: int32
//
// responses:
//
//  '200':
//   description: Success
//   schema:
//    "$ref": "#/definitions/Foundation"
//  '400':
//   description: Bad Request
//  '401':
//   description: Unauthorized
//  '404':
//   description: Not Found
//  '409':
//   description: Conflict
func Update(log *slog.Logger, ctrl IUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var err error
		var response fndResp.BaseResponse
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
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
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
