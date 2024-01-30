package users

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IByLoginGetter interface {
	GetByLogin(login string) (ents.UserTransfer, error)
}

func GetByLogin(log *slog.Logger, ctrl IByLoginGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.BaseResponse
		var user ents.UserTransfer
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		const op = "handlers.user.getByLogin"
		log = log.With(
			slog.String("operation", op),
		)
		login := mux.Vars(r)["login"]
		user, err = ctrl.GetByLogin(login)
		if err != nil {
			return
		}
		response = usrResp.BaseResponse{User: user}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)

	}
}
