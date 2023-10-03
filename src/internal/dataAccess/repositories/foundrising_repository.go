package repositories

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type IFoundrisingRepository interface {
	Insert(F ents.Foundrising) error
	Update(F ents.Foundrising) error
	Delete(F ents.Foundrising) error
	Select() ([]ents.Foundrising, bool)
	SelectById(id uint64) (ents.Foundrising, bool)
	SelectByFoundId(id uint64) ([]ents.Foundrising, bool)
	SelectByCreateDate(date string) ([]ents.Foundrising, bool)
	SelectByFoundIdActive(id uint64) ([]ents.Foundrising, bool)
	SelectByCloseDate(date string) ([]ents.Foundrising, bool)
	SelectByIdAndFoundId(id_ uint64, found_id_ uint64) (ents.Foundrising, bool)
	//GetDB() *gorm.DB
}
