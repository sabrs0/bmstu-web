package foundrisings

import (
	"encoding/json"
	"log/slog"
	"net/http"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	fndResp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response/foundrisings"
)

type IAdder interface {
	Add(params ents.FoundrisingAdd) (ents.Foundrising, error)
}

// swagger:route POST /foundrisings Foundrising FoundrisingsPost
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
//	   200: FoundrisingsBaseResponse
//	   400: ValidateError
//	   401: ValidateError
//	   409: ValidateError
func Add(log *slog.Logger, ctrl IAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response fndResp.BaseResponse
		const op = "handlers.foundrisings.add"
		log = log.With(
			slog.String("operation", op),
		)
		defer func() {
			if err != nil {
				log.Error(err.Error())
				resp.ErrWrapper(w, &response, err)
			}

		}()
		var params ents.FoundrisingAdd
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			return
		}
		foundrising, err := ctrl.Add(params)
		if err != nil {
			return
		}
		log.Info("Success")
		response = fndResp.BaseResponse{Foundrising: foundrising}
		resp.JSONRender(w, http.StatusOK, &response)
	}
}
