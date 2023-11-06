package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IUpdater interface {
	Update(id string, params ents.UserAdd) (ents.User, error)
}

// swagger:operation PUT /users/{id} UsersUpdate
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
//	    "$ref": "#/definitions/UserAdd"
//
// responses:
//
//	'200':
//	  description: Success
//	  schema:
//	    "$ref": "#/definitions/User"
func Update(log *slog.Logger, ctrl IUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.BaseResponse
		const op = "handlers.users.update"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		var params ents.UserAdd
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		id := mux.Vars(r)["id"]
		user, err := ctrl.Update(id, params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = usrResp.BaseResponse{User: user}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
