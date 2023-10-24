package foundrisings

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IGetter interface {
	GetAll() ([]ents.Foundrising, error)
}

func GetAll(log *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.GetAllResponse
		const op = "handlers.foundrisings.getAll"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			resp.ErrWrapper(log, w, response, err)
		}()
		var foundrisings []ents.Foundrising
		foundrisings, err = ctrl.GetAll()
		if err != nil {
			return
		}
		response = fndResp.GetAllResponse{
			Foundrisings: foundrisings,
		}
		log.Info("Success")
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
