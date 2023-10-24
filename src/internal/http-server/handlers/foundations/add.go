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

func Add(logger *slog.Logger, ctrl IAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.foundations.Add"
		logger = logger.With(
			slog.String("Operation", op),
		)
		var params ents.FoundationAdd
		var response fndResp.BaseResponse
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			resp.JSONRender(w, http.StatusBadRequest, response)
			logger.Error(err.Error())
			return
		}
		logger.Info("Request body decoded")
		foundation, err := ctrl.Add(params)
		if err != nil {
			resp.JSONRender(w, http.StatusConflict, resp.Response{
				Status: "Error",
				Error:  "Failed add foundation" + err.Error(),
			})
			logger.Error(err.Error())
			return
		}

		logger.Info("Successfully added foundation")
		response = fndResp.BaseResponse{Foundation: foundation}
		resp.JSONRender(w, http.StatusOK, response)
	}

}
