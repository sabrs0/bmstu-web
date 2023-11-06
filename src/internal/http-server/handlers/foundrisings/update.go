package foundrisings

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IUpdater interface {
	Update(id string, params ents.FoundrisingPut) (ents.Foundrising, error)
}

// swagger:operation PUT /foundrisings/{id} FoundrisingsUpdate
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
// requestBody:
//
//	schema:
//	    "$ref": "#/definitions/FoundrisingPut"
//
// responses:
//
//	'200':
//	  description: Success
//	  schema:
//	    "$ref": "#/definitions/Foundrising"
func Update(log *slog.Logger, ctrl IUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		const op = "handlers.foundrisings.update"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		var params ents.FoundrisingPut
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		err = params.ValidatePut()
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		foundrising, err := ctrl.Update(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundrising: foundrising}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}

func UpdateOptional(log *slog.Logger, ctrl IUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		const op = "handlers.foundrisings.update"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)

			}
		}()
		var params ents.FoundrisingPut
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		err = params.ValidatePatch()
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		foundrising, err := ctrl.Update(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundrising: foundrising}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
