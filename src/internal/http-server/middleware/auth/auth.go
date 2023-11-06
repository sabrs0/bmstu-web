package auth

import (
	"log"
	"net/http"

	resp "github.com/sabrs0/bmstu-web/src/internal/lib/api/response"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type RequestValidator interface {
	ValidateRequest(r *http.Request, roleToAccept []string) error
}

func New(roles []string, validator RequestValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := validator.ValidateRequest(r, roles)
			if err != nil {
				log.Printf("AUTH MW ERROR %s\n", err.Error())
				resp.ErrWrapper(w, nil, my_errors.ErrorAuth)
				return
			}

			log.Printf("TOKEN IS VALID\n")
			next.ServeHTTP(w, r)
		})
	}
}
