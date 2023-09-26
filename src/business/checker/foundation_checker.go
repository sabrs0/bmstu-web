package checker

import (
	"fmt"

	"github.com/sabrs0/bmstu-web/src/business/entities"
	"github.com/sabrs0/bmstu-web/src/my_errors"
)

type FoundationDonateParams struct {
	Sum_of_money        float64
	IsClosedFoundrising bool
}
type FoundationMainParams struct {
	Login    string
	Name     string
	Password string
	Country  string
}

func (UP *FoundationMainParams) CheckParams() error {
	if CheckLogin(UP.Login) {
		return fmt.Errorf("error incorrect Foundationname")
	}

	return CheckCountry(UP.Country)
}
func NewFoundationMainParams(login string, passw string, name string, country string) FoundationMainParams {
	return FoundationMainParams{Login: login, Name: name, Password: passw, Country: country}
}
func NewFoundationDonateParams(sum float64, isClosed bool) FoundationDonateParams {
	return FoundationDonateParams{Sum_of_money: sum, IsClosedFoundrising: isClosed}
}
func CheckCountry(c string) error {
	var flag bool = false
	for _, cnt := range entities.Countries {
		if c == cnt {
			flag = true
		}
	}
	if !flag {
		return fmt.Errorf(my_errors.ErrCountry)
	}
	return nil

}
