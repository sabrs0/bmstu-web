package foundations

import (
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IGetter interface {
	GetAll() ([]ents.FoundationTransfer, error)
}

// swagger:route GET /foundations Foundation FoundationsGet
//
//		 Consumes:
//		 - application/json
//
//		 Produces:
//		 - application/json
//
//		 Schemes: http
//
//
//
//		 Responses:
//	  		default: ValidateError
//		    200: FoundationsGetAllResponse
func GetAll(logger *slog.Logger, ctrl IGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.GetAllResponse
		const op = "handlers.foundations.getAll"
		logger = slog.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				logger.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		var foundations []ents.FoundationTransfer
		foundations, err = ctrl.GetAll()
		if err != nil {
			return
		}
		logger.Info("Success")
		response.Foundations = foundations
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
