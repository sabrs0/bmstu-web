package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IAdder interface {
	Add(params ents.UserAdd) (ents.User, error)
}

// swagger:route POST /users User UsersPost
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
//  Security:
//	  bearerAuth:
//
//  Responses:
//    default: ValidateError
//    200: UsersBaseResponse
//    400: ValidateError
//    401: ValidateError
//    409: ValidateError

/*
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - text/plain
// requestBody:
//
//	schema:
//	 "$ref": "#/definitions/UserAdd"
//
// responses:
//
//	'200':
//	 description: Success
//	 schema:
//	   "$ref": "#/definitions/User"
//	'400':
//	 description: Bad Request
//	'401':
//	 description: Unauthorized
//	'409':
//	 description: Conflict
*/
func Add(log *slog.Logger, ctrl IAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.BaseResponse
		const op = "handlers.users.add"
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
		user, err := ctrl.Add(params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = usrResp.BaseResponse{User: user}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
