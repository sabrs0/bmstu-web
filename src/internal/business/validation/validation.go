package validation

import (
	"fmt"

	"github.com/sabrs0/bmstu-web/src/internal/business/entities"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

func CheckLogin(name string) bool {
	for _, c := range name {
		if (c < 'a' || c > 'z') && !(c >= 'A' && c <= 'Z') && !(c >= '0' && c <= '9') {
			return true
		}
	}
	return false
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
func CheckMoneyFormat(money string) error {
	for i, c := range money {
		if c == '.' {
			if i < len(money)-1 && len(money)-1-i < 3 && i > 0 {
				for j := i + 1; j < len(money); j++ {
					if money[j] < '0' || money[j] > '9' {
						return fmt.Errorf(my_errors.ErrMoney)
					}
				}
			} else {
				return fmt.Errorf(my_errors.ErrMoney)
			}
		}
	}
	return nil
}
