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
func (UC *FoundationController) GetAll() ([]ents.FoundationTransfer, error) {
	foundartions, err := UC.FS.GetAll()
	if err != nil {
		return nil, err
	}
	ft := []ents.FoundationTransfer{}
	for _, f := range foundartions {
		ft = append(ft, ents.NewFoundationTransfer(f))
	}
	return ft, nil
}
func (UC *FoundationController) GetAllFoundrisings(id string) ([]ents.FoundrisingTransfer, error) {
	foundrisings, err := UC.FgS.GetByFoundId(id)
	if err != nil {
		return nil, err
	}
	ft := []ents.FoundrisingTransfer{}
	for _, f := range foundrisings {
		ft = append(ft, ents.NewFoundrisingTransfer(f))
	}
	return ft, nil
}
func (UC *FoundationController) GetByID(id_ string) (ents.FoundationTransfer, error) {
	f, err := UC.FS.GetById(id_)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil
}
func (UC *FoundationController) GetExtByID(id_ string) (ents.FoundationTransferExtended, error) {
	f, err := UC.FS.GetById(id_)
	if err != nil {
		return ents.FoundationTransferExtended{}, err
	}
	return ents.NewFoundationTransferExtended(f), nil
}
func (UC *FoundationController) GetExtByLogin(login string) (ents.FoundationTransferExtended, error) {
	f, err := UC.FS.GetByLogin(login)
	if err != nil {
		return ents.FoundationTransferExtended{}, err
	}
	return ents.NewFoundationTransferExtended(f), nil
}
func (UC *FoundationController) GetByName(name string) (ents.FoundationTransfer, error) {
	f, err := UC.FS.GetByName(name)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil
}
func (UC *FoundationController) GetByLogin(login string) (ents.FoundationTransfer, error) {
	f, err := UC.FS.GetByLogin(login)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil
}

func (UC *FoundationController) GetByCountry(country string) ([]ents.FoundationTransfer, error) {
	foundartions, err := UC.FS.GetByCountry(country)
	if err != nil {
		return nil, err
	}
	ft := []ents.FoundationTransfer{}
	for _, f := range foundartions {
		ft = append(ft, ents.NewFoundationTransfer(f))
	}
	return ft, nil
}

func (UC *FoundationController) Add(params ents.FoundationAdd) (ents.FoundationTransfer, error) {
	foundations_by_country, err := UC.FS.GetByCountry(params.Country)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	for _, f := range foundations_by_country {
		if f.Name == params.Name {
			return ents.FoundationTransfer{}, fmt.Errorf("в данной стране уже есть фонд с таким названием")
		}
	}
	f, err := UC.FS.Add(params)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil

}

func (UC *FoundationController) Delete(id string) (ents.FoundationTransfer, error) { //это делают триггеры (каскадное удаление)

	f, err := UC.FS.Delete(id)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil

}
func (UC *FoundationController) Update(id string, params ents.FoundationAdd) (ents.FoundationTransfer, error) {
	var foundation ents.Foundation
	foundation, _ = UC.FS.GetById(id) //US.GetByID(id)
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
	f, err := UC.FS.Update(id, params)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(f), nil

}
func (UC *FoundationController) AcceptDonate(id string, sum float64) error {
	return UC.FS.AcceptDonate(id, sum)
}
func (UC *FoundationController) DonateToFoundrising(foundationID string,
	params ents.FoundationDonate) (ents.TransactionTransfer, error) {
	var Transaction ents.Transaction
	err := validation.CheckMoneyFormat(params.Sum_of_Money)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	reqSum, err := strconv.ParseFloat(params.Sum_of_Money, 64)
	if err != nil || params.Sum_of_Money == "NaN" || params.Sum_of_Money == "Infinity" || params.Sum_of_Money == "Inf" {
		return ents.TransactionTransfer{}, err
	}
	U, err := UC.FS.GetById(foundationID) //UC.GetByID(foundationID)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	if U.Fund_balance < reqSum {
		return ents.TransactionTransfer{}, fmt.Errorf("недостаточно средств ")
	}
	foundrising, err := UC.FgS.GetById(params.Foundrising_id)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	if foundrising.Closing_date.Valid {
		return ents.TransactionTransfer{}, fmt.Errorf("данный сбор уже закрыт")
	}
	err = UC.FS.Donate(&U, reqSum)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	var remainder float64
	remainder, err = UC.FgS.AcceptDonate(params.Foundrising_id, reqSum, false)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	sid, err := strconv.Atoi(params.Foundrising_id)
	foundrising_id := uint64(sid)
	if err != nil {
		return ents.TransactionTransfer{}, err
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
		return ents.TransactionTransfer{}, err
	}
	if remainder > 1e-9 {
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
		return ents.TransactionTransfer{}, err
	} else if remainder <= 1e-9 && remainder > -1.0 {
		U.ClosedFoundrisingAmount += 1
		if U.CurFoudrisingAmount > 0 {
			U.CurFoudrisingAmount -= 1
		}
		//это кошмарище:
		_, err := UC.FS.GetRepo().Update(U)
		return ents.TransactionTransfer{}, err
	}
	return ents.NewTransactionTransfer(Transaction), nil
}

func (UC *FoundationController) ReplenishBalance(id string, params ents.FoundationReplenish) (ents.FoundationTransfer, error) {
	var err error
	var reqSum float64

	reqSum, err = strconv.ParseFloat(params.Sum_of_Money, 64)
	if err != nil || params.Sum_of_Money == "NaN" || params.Sum_of_Money == "Infinity" || params.Sum_of_Money == "Inf" {
		return ents.FoundationTransfer{}, err
	}
	err = validation.CheckMoneyFormat(params.Sum_of_Money)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	U, err := UC.FS.GetById(id) //UC.GetByID(foundationID)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	if reqSum > 50000.00 {
		return ents.FoundationTransfer{}, fmt.Errorf("введенная сумма превышается 50 000")
	}

	_, err = UC.FS.ReplenishBalance(&U, reqSum)
	if err != nil {
		return ents.FoundationTransfer{}, err
	}
	return ents.NewFoundationTransfer(U), nil
}
