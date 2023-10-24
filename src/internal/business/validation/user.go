package validation

import (
	"fmt"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

func CheckUserAddParams(UP ents.UserAdd) error {
	if CheckLogin(UP.Login) {
		return fmt.Errorf("error incorrect username")
	}
	return nil
}
