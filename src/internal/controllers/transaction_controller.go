package controllers

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"

	"strconv"
)

type TransactionController struct {
	TS  servs.TransactionService
	FS  servs.FoundationService
	FgS servs.FoundrisingService
	US  servs.UserService // индусский код
}

func NewTransactionController(FdR servs.IFoundationRepository, FgR servs.IFoundrisingRepository,
	UR servs.IUserRepository, TR servs.ITransactionRepository) *TransactionController {
	return &TransactionController{
		TS:  *servs.NewTransactionService(TR),
		FS:  *servs.NewFoundationService(FdR),
		FgS: *servs.NewFoundrisingService(FgR),
		US:  *servs.NewUserService(UR),
	}
}

func (TC *TransactionController) GetAll() ([]ents.Transaction, error) {
	return TC.TS.GetAll()
}
func (TC *TransactionController) GetByID(id_ string) (ents.Transaction, error) {
	return TC.TS.GetById(id_)
}
func (TC *TransactionController) GetFromId(type_ string, id_ string) ([]ents.Transaction, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err != nil {
		return nil, err
	}
	Transactions, err = TC.TS.GetFromId(booltype, id_, &TC.FS, &TC.US)
	return Transactions, err
}
func (TC *TransactionController) GetToId(type_ string, id_ string) ([]ents.Transaction, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err != nil {
		return nil, err
	}
	Transactions, err = TC.TS.GetToId(booltype, id_, &TC.FS, &TC.FgS)

	return Transactions, err
}

func (TC *TransactionController) Delete(id string) (ents.Transaction, error) {
	return TC.TS.Delete(id)

}

func (TC *TransactionController) GetFoundrisingPhilantropIds(id_ string) ([]uint64, error) {
	return TC.TS.GetFoundrisingPhilantropIds(id_, &TC.FgS)
}
