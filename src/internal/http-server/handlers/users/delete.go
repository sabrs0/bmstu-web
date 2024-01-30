package users

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IDeleter interface {
	Delete(id string) (ents.UserTransfer, error)
}

// swagger:route DELETE /users/{id} User UsersDelete
//
//	 Consumes:
//	 - application/json
//	 - text/plain
//
//	 Produces:
//	 - application/json
//	 - text/plain
//
//	 Schemes: http
//
//
//	 Security:
//		  bearerAuth:
//
//	 Parameters:
//	      + name: id
//	        in: path
//	        required: true
//	        type: integer
//	        format: int64
//	 Responses:
//	   default: ValidateError
//	   200: UsersBaseResponse
//	   401: ValidateError
//	   404: ValidateError
func Delete(log *slog.Logger, ctrl IDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.BaseResponse
		const op = "handlers.users.delete"
		var user ents.UserTransfer
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
		user, err = ctrl.Delete(id)
		if err != nil {
			return
		}
		log.Info("Success")
		response = usrResp.BaseResponse{User: user}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
