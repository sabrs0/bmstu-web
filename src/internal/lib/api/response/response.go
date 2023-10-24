package response

import (
	"encoding/json"
	"net/http"
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
