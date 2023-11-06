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

// swagger:operation POST /foundations Foundation FoundationsPost
//
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
//	  "$ref": "#/definitions/FoundationAdd"
//
// responses:
//
//	'200':
//	 description: Success
//	 schema:
//	   "$ref": "#/definitions/Foundation"
//	'400':
//	 description: Bad Request
//	'401':
//	 description: Unauthorized
//	'409':
//	 description: Conflict
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
