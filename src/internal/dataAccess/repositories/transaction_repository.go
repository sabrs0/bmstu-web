package repositories

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type ITransactionRepository interface {
	Insert(T ents.Transaction) error
	Delete(T ents.Transaction) error
	Select() ([]ents.Transaction, bool)
	SelectFrom(type_ bool, id uint64) ([]ents.Transaction, bool)
	SelectTo(type_ bool, id uint64) ([]ents.Transaction, bool)
	SelectById(id uint64) (ents.Transaction, bool)
	SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, bool)
	//GetDB() *gorm.DB
}
