package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

// swagger:response ValidateError
type ErrResponse struct {
	//in: body
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(msg string) ErrResponse {
	return ErrResponse{
		Error: msg,
	}
}

func ErrWrapper(w http.ResponseWriter, resp any, err error) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusBadRequest
	if err == my_errors.ErrorNotExists {
		status = http.StatusNotFound
	}
	if err == my_errors.ErrorConflict {
		status = http.StatusConflict
	}
	if err == my_errors.ErrorAuth {
		status = http.StatusUnauthorized
	}

	//http.Error(w, err.Error(), status)
	w.WriteHeader(status)
}
func JSONRender(w http.ResponseWriter, status int, jsonData any) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(jsonData)
	fmt.Println("JSONNED DATA IS: ", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
