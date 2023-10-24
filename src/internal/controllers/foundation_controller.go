package controllers

import (
	"fmt"
	"strconv"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/internal/business/validation"
)

type FoundationController struct {
	FS  servs.FoundationService
	TS  servs.TransactionService
	FgS servs.FoundrisingService
}

func NewFoundationController(FR servs.IFoundationRepository, TR servs.ITransactionRepository,
	FgR servs.IFoundrisingRepository) *FoundationController {
	return &FoundationController{
		FS:  *servs.NewFoundationService(FR),
		TS:  *servs.NewTransactionService(TR),
		FgS: *servs.NewFoundrisingService(FgR),
	}
}
func (UC *FoundationController) GetAll() ([]ents.Foundation, error) {
	return UC.FS.GetAll()
}
func (UC *FoundationController) GetAllFoundrisings(id string) ([]ents.Foundrising, error) {
	return UC.FgS.GetByFoundId(id)
}
func (UC *FoundationController) GetByID(id_ string) (ents.Foundation, error) {
	return UC.FS.GetById(id_)
}
func (UC *FoundationController) GetByName(name string) (ents.Foundation, error) {
	return UC.FS.GetByName(name)
}
func (UC *FoundationController) GetByLogin(login string) (ents.Foundation, error) {
	return UC.FS.GetByLogin(login)
}

func (UC *FoundationController) GetByCountry(country string) ([]ents.Foundation, error) {
	return UC.FS.GetByCountry(country)
}

func (UC *FoundationController) Add(params ents.FoundationAdd) (ents.Foundation, error) {
	foundations_by_country, err := UC.FS.GetByCountry(params.Country)
	if err != nil {
		return ents.Foundation{}, err
	}
	for _, f := range foundations_by_country {
		if f.Name == params.Name {
			return ents.Foundation{}, fmt.Errorf("в данной стране уже есть фонд с таким названием")
		}
	}
	return UC.FS.Add(params)

}

func (UC *FoundationController) Delete(id string) (ents.Foundation, error) { //это делают триггеры (каскадное удаление)

	return UC.FS.Delete(id)

}
func (UC *FoundationController) Update(id string, params ents.FoundationAdd) (ents.Foundation, error) {
	var foundation ents.Foundation
	foundation, _ = UC.GetByID(id)
	if params.Login == "" {
		params.Login = foundation.Login
	}
	if params.Password == "" {
		params.Password = foundation.Password
	}
	if params.Name == "" {
		params.Name = foundation.Name
	}
	if params.Country == "" {
		params.Country = foundation.Country
	}
	return UC.FS.Update(id, params)

}
func (UC *FoundationController) AcceptDonate(id string, sum float64) error {
	return UC.FS.AcceptDonate(id, sum)
}
func (UC *FoundationController) DonateToFoundrising(foundationID string,
	params ents.FoundationDonate) (ents.Transaction, error) {
	var Transaction ents.Transaction
	err := validation.CheckMoneyFormat(params.Sum_of_Money)
	if err != nil {
		return ents.Transaction{}, err
	}
	reqSum, err := strconv.ParseFloat(params.Sum_of_Money, 64)
	if err != nil {
		return ents.Transaction{}, err
	}
	U, err := UC.GetByID(foundationID)
	if err != nil {
		return ents.Transaction{}, err
	}
	if U.Fund_balance < reqSum {
		return ents.Transaction{}, fmt.Errorf("недостаточно средств ")
	}
	foundrising, err := UC.FgS.GetById(params.Foundrising_id)
	if err != nil {
		return ents.Transaction{}, err
	}
	if foundrising.Closing_date.Valid {
		return ents.Transaction{}, fmt.Errorf("данный сбор уже закрыт")
	}
	err = UC.FS.Donate(&U, reqSum)
	if err != nil {
		return ents.Transaction{}, err
	}
	var remainder float64
	remainder, err = UC.FgS.AcceptDonate(params.Foundrising_id, reqSum, false)
	if err != nil {
		return ents.Transaction{}, err
	}
	sid, err := strconv.Atoi(params.Foundrising_id)
	foundrising_id := uint64(sid)
	if err != nil {
		return ents.Transaction{}, err
	}
	TP := ents.TransactionAdd{
		From_essence_type: ents.FROM_FOUNDATION,
		From_id:           U.Id,
		To_essence_type:   ents.TO_FOUNDRISING,
		Sum_of_money:      reqSum,
		Comment:           params.Comm,
		To_id:             foundrising_id,
	}
	Transaction, err = UC.TS.Add(TP)
	if err != nil {
		return ents.Transaction{}, err
	}
	if remainder > 0.0 {
		foundrising, _ := UC.FgS.GetById(params.Foundrising_id)
		found_id := foundrising.Found_id
		TP := ents.TransactionAdd{
			From_essence_type: ents.FROM_FOUNDATION,
			From_id:           U.Id,
			To_essence_type:   ents.TO_FOUNDATION,
			Sum_of_money:      remainder,
			Comment:           "returning the remain to myself",
			To_id:             found_id,
		}
		_, err = UC.TS.Add(TP)
		return ents.Transaction{}, err
	} else if remainder <= 1e-9 {
		U.ClosedFoundrisingAmount += 1
		if U.CurFoudrisingAmount > 0 {
			U.CurFoudrisingAmount -= 1
		}
		//это кошмарище:
		_, err := UC.FS.GetRepo().Update(U)
		return ents.Transaction{}, err
	}
	return Transaction, nil
}

func (UC *FoundationController) ReplenishBalance(sum string, U *ents.Foundation) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	if err != nil {
		return err
	}
	err = validation.CheckMoneyFormat(sum)
	if err != nil {
		return err
	}
	if reqSum > 50000.00 {
		return fmt.Errorf("введенная сумма превышается 50 000")
	}
	return UC.FS.ReplenishBalance(U, reqSum)

}
