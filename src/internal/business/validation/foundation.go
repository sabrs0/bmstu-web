package validation

import (
	"fmt"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

func CheckFoundationAddParams(UP ents.FoundationAdd) error {
	if CheckLogin(UP.Login) {
		return fmt.Errorf("error incorrect Foundationname")
	}

	return CheckCountry(UP.Country)
}
