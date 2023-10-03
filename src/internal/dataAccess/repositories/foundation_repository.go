package repositories

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	"gorm.io/gorm"
)

type IFoundationRepository interface {
	Insert(F ents.Foundation) error
	Update(F ents.Foundation) error
	Delete(F ents.Foundation) error
	Select() ([]ents.Foundation, bool)
	SelectById(id uint64) (ents.Foundation, bool)
	SelectByLogin(name string) (ents.Foundation, bool)
	SelectByName(name string) (ents.Foundation, bool)
	SelectByCountry(country string) ([]ents.Foundation, bool)
	GetDB() *gorm.DB
}
