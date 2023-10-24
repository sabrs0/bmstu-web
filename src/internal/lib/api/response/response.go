package response

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(msg string) Response {
	return Response{
		Error: msg,
	}
}
func ErrWrapper(log *slog.Logger, w http.ResponseWriter, resp any, err error) {
	status := http.StatusBadRequest
	if err == my_errors.ErrorNotExists {
		status = http.StatusNotFound
	}
	if err == my_errors.ErrorConflict {
		status = http.StatusConflict
	}
	log.Error(err.Error())
	JSONRender(w, status, resp)
}
func JSONRender(w http.ResponseWriter, status int, jsonData any) {
	data, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
