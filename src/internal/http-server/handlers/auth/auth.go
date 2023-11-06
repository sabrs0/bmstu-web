package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	auth "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	authResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/auth"
)

type IAuther interface {
	Login(params auth.Params) (string, error)
}

// swagger:operation POST /login Login Login
//
//
//
//
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - text/plain
// requestBody:
//    schema:
//        "$ref": "#/definitions/LoginParams"
// responses:
// '200':
//   description: Success
//   schema:
//    "$ref": "#/definitions/Token"
// '400':
//   description: Bad Request
// '404':
//   description: Not Found

func Login(log *slog.Logger, ctrl IAuther) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response authResp.TokenResponse
		const op = "handlers.auth.login"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		var params auth.Params
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		token, err := ctrl.Login(params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = authResp.TokenResponse{Token: token}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
