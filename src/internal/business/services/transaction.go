package services

import (
	"fmt"
	"strconv"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type ITransactionRepository interface {
	Insert(T ents.Transaction) (ents.Transaction, error)
	Delete(T ents.Transaction) (ents.Transaction, error)
	Select() ([]ents.Transaction, error)
	SelectFrom(type_ bool, id uint64) ([]ents.Transaction, error)
	SelectTo(type_ bool, id uint64) ([]ents.Transaction, error)
	SelectById(id uint64) (ents.Transaction, error)
	SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, error)
	//GetDB() *gorm.DB
}

type ExistByIdChecker interface {
	ExistsById(id uint64) bool
}

type ITransactionService interface {
	Add(TPars ents.TransactionAdd) (ents.Transaction, error)
	Delete(id_ string) (ents.Transaction, error)
	GetAll() ([]ents.Transaction, error)
	GetById(id_ string) (ents.Transaction, error)
	GetFromId(type_ bool, id_ string, FndS ExistByIdChecker,
		US ExistByIdChecker) ([]ents.Transaction, error)
	GetFoundrisingPhilantropIds(id_ string, FndgS ExistByIdChecker) ([]uint64, error)
	GetToId(type_ bool, id_ string, FndS ExistByIdChecker,
		FndgS ExistByIdChecker) ([]ents.Transaction, error)

	ExistsById(id uint64) bool
	GetRepo() ITransactionRepository
}

type TransactionService struct {
	TR ITransactionRepository
}

func NewTransactionService(frepo ITransactionRepository) *TransactionService {
	return &TransactionService{TR: frepo}
}
func (TS *TransactionService) GetRepo() ITransactionRepository {
	return TS.TR
}
func (FS *TransactionService) ExistsById(id uint64) bool {
	_, result := FS.TR.SelectById(id)
	return (result == nil)
}

func (FS *TransactionService) Add(TPars ents.TransactionAdd) (ents.Transaction, error) {

	var U ents.Transaction = ents.NewTransaction()
	U.SetFromId(TPars.From_id)
	U.SetFromType(TPars.From_essence_type)
	U.SetComment(TPars.Comment)
	U.SetSumOfMoney(TPars.Sum_of_money)
	U.SetToId(TPars.To_id)
	U.SetToType(TPars.To_essence_type)

	return FS.TR.Insert(U)
}
func (FS *TransactionService) Delete(id_ string) (ents.Transaction, error) {
	var errGet error
	var U ents.Transaction
	U, errGet = FS.GetById(id_)
	if errGet != nil {
		return ents.Transaction{}, errGet
	}
	return FS.TR.Delete(U)
}
func (FS *TransactionService) GetAll() ([]ents.Transaction, error) {
	Transactions, err := FS.TR.Select()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список всех транзакций: %s", err)
	}
	return Transactions, nil
}
func (FS *TransactionService) GetById(id_ string) (ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return ents.Transaction{}, fmt.Errorf("некорректный id")
	}
	if !FS.ExistsById(id) {
		return ents.Transaction{}, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
	}
	U, err := FS.TR.SelectById(id)
	if err != nil {
		return U, fmt.Errorf("не удалось получить транзакцию по id: %s", err)
	}

	return U, nil
}

// в списке аргументов полный кринжжжж
func (TS *TransactionService) GetFromId(type_ bool, id_ string, FndS ExistByIdChecker,
	US ExistByIdChecker) ([]ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return nil, fmt.Errorf("некорректный id")
	}
	var U []ents.Transaction
	if type_ == ents.FROM_USER {
		if !US.ExistsById(id) {
			return nil, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
		}
		U, err = TS.TR.SelectFrom(type_, id)
		if err != nil {
			return nil, fmt.Errorf("не удалось получить транзакцию по id отправителя: %s", err)
		}
	} else {
		if !FndS.ExistsById(id) {
			return nil, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
		}
		U, err = TS.TR.SelectFrom(type_, id)
		if err != nil {
			return U, fmt.Errorf("не удалось получить транзакцию по id отправителя: %s", err)
		}
	}
	return U, nil
}

func (TS *TransactionService) GetToId(type_ bool, id_ string, FndS ExistByIdChecker,
	FndgS ExistByIdChecker) ([]ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U []ents.Transaction
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	}
	if type_ == ents.TO_FOUNDATION {
		if !FndS.ExistsById(id) {
			return U, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
		}
		U, err = TS.TR.SelectTo(type_, id)
		if err != nil {
			return U, fmt.Errorf("не удалось получить транзакцию по id получателя: %s", err)
		}
	} else {
		if !FndgS.ExistsById(id) {
			return U, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
		}
		U, err = TS.TR.SelectTo(type_, id)
		if err != nil {
			return U, fmt.Errorf("не удалось получить транзакцию по id получателя: %s", err)
		}
	}
	return U, nil
}

func (TS *TransactionService) GetFoundrisingPhilantropIds(id_ string, FndgS ExistByIdChecker) ([]uint64, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var IDs []uint64
	if err != nil {
		return IDs, fmt.Errorf("некорректный id")
	}
	if !FndgS.ExistsById(id) {
		return IDs, fmt.Errorf("несуществующий id: %w", my_errors.ErrorNotExists)
	}
	IDs, err = TS.TR.SelectFoundrisingPhilantropIds(id)
	if err != nil {
		return IDs, fmt.Errorf("не удалось получить id филантропов : %s", err)
	}
	return IDs, nil

}
