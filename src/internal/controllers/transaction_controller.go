package controllers

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	repos "github.com/sabrs0/bmstu-web/src/internal/dataAccess/repositories/gorm"

	"strconv"

	"gorm.io/gorm"
)

type TransactionController struct {
	TS  servs.ITransactionService
	FS  servs.IFoundationService
	FgS servs.IFoundrisingService
	US  servs.IUserService
}

func NewTransactionController(db *gorm.DB, fs servs.IFoundationService, fgs servs.IFoundrisingService,
	us servs.IUserService) *TransactionController {
	TR := repos.NewTransactionRepository(db)
	tS := servs.NewTransactionService(TR)
	return &TransactionController{
		TS:  tS,
		FS:  fs,
		FgS: fgs,
		US:  us,
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
	if err == nil {
		Transactions, err = TC.TS.GetFromId(booltype, id_, TC.FS, TC.US)
	}
	return Transactions, err
}
func (TC *TransactionController) GetToId(type_ string, id_ string) ([]ents.Transaction, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err == nil {
		Transactions, err = TC.TS.GetToId(booltype, id_, TC.FS, TC.FgS)
	}
	return Transactions, err
}

func (TC *TransactionController) Delete(id string) error {
	return TC.TS.Delete(id)

}

func (TC *TransactionController) GetFoundrisingPhilantropIds(id_ string) ([]uint64, error) {
	return TC.TS.GetFoundrisingPhilantropIds(id_, TC.FgS)
}
