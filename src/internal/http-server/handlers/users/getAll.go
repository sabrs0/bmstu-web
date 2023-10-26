package users

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	usrResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/users"
)

type IGetter interface {
	GetAll() ([]ents.User, error)
}

func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response usrResp.GetAllResponse
		const op = "handlers.users.getAll"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				resp.ErrWrapper(log, w, &response, err)
			}

		}()
		users, err := ctrl.GetAll()
		if err != nil {
			return
		}
		response = usrResp.GetAllResponse{
			Users: users,
		}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
