package foundations

import (
	"encoding/json"
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundations"
)

type IAdder interface {
	Add(params ents.FoundationAdd) (ents.Foundation, error)
}

// swagger:route POST /foundations Foundation FoundationsPost
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Schemes: http
//
//
//	 Security:
//		  bearerAuth:
//
//	 Responses:
//	   default: ValidateError
//	   200: FoundationsBaseResponse
//	   400: ValidateError
//	   401: ValidateError
//	   409: ValidateError
func Add(logger *slog.Logger, ctrl IAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.foundations.Add"
		logger = logger.With(
			slog.String("Operation", op),
		)
		var err error
		var params ents.FoundationAdd
		var response fndResp.BaseResponse
		defer func() {
			if err != nil {
				logger.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {

			return
		}
		logger.Info("Request body decoded")
		foundation, err := ctrl.Add(params)
		if err != nil {

			return
		}

		logger.Info("Success")
		response = fndResp.BaseResponse{Foundation: foundation}
		resp.JSONRender(w, http.StatusOK, &response)
	}

}
