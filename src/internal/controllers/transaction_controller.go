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

func (TC *TransactionController) GetAll() ([]ents.TransactionTransfer, error) {
	transactions, err := TC.TS.GetAll()
	if err != nil {
		return nil, err
	}
	var transactionsTransfer []ents.TransactionTransfer
	for _, tr := range transactions {
		transactionsTransfer = append(transactionsTransfer, ents.NewTransactionTransfer(tr))
	}
	return transactionsTransfer, nil
}
func (TC *TransactionController) GetByID(id_ string) (ents.TransactionTransfer, error) {
	tr, err := TC.TS.GetById(id_)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	return ents.NewTransactionTransfer(tr), nil
}
func (TC *TransactionController) GetFromId(type_ string, id_ string) ([]ents.TransactionTransfer, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err != nil {
		return nil, err
	}
	Transactions, err = TC.TS.GetFromId(booltype, id_, &TC.FS, &TC.US)
	if err != nil {
		return nil, err
	}
	var transactionsTransfer []ents.TransactionTransfer
	for _, tr := range Transactions {
		transactionsTransfer = append(transactionsTransfer, ents.NewTransactionTransfer(tr))
	}
	return transactionsTransfer, nil
}
func (TC *TransactionController) GetToId(type_ string, id_ string) ([]ents.TransactionTransfer, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err != nil {
		return nil, err
	}
	Transactions, err = TC.TS.GetToId(booltype, id_, &TC.FS, &TC.FgS)
	if err != nil {
		return nil, err
	}
	var transactionsTransfer []ents.TransactionTransfer
	for _, tr := range Transactions {
		transactionsTransfer = append(transactionsTransfer, ents.NewTransactionTransfer(tr))
	}
	return transactionsTransfer, nil
}

func (TC *TransactionController) Delete(id string) (ents.TransactionTransfer, error) {
	tr, err := TC.TS.Delete(id)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	return ents.NewTransactionTransfer(tr), nil
}

func (TC *TransactionController) GetFoundrisingPhilantropIds(id_ string) ([]uint64, error) {
	return TC.TS.GetFoundrisingPhilantropIds(id_, &TC.FgS)
}
