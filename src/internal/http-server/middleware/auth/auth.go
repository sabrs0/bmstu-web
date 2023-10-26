package auth

import (
	"fmt"
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

				resp.JSONRender(w, http.StatusUnauthorized, resp.Response{
					Error: fmt.Errorf("%w: %w", my_errors.ErrorAuth, err).Error(),
				})
			}
			next.ServeHTTP(w, r)
		})
	}
}
