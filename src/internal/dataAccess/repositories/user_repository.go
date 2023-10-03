package repositories

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type IUserRepository interface {
	Insert(U ents.User) error
	Update(U ents.User) error
	Delete(U ents.User) error
	Select() ([]ents.User, bool)
	SelectById(id uint64) (ents.User, bool)
	SelectByLogin(name string) (ents.User, bool)
	//GetDB() *gorm.DB
}
