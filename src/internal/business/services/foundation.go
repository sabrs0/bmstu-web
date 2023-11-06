package services

import (
	"fmt"
	"strconv"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	"github.com/sabrs0/bmstu-web/src/internal/business/validation"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IFoundationRepository interface {
	Insert(F ents.Foundation) (ents.Foundation, error)
	Update(F ents.Foundation) (ents.Foundation, error)
	Delete(F ents.Foundation) (ents.Foundation, error)
	Select() ([]ents.Foundation, error)
	SelectById(id uint64) (ents.Foundation, error)
	SelectByLogin(name string) (ents.Foundation, error)
	SelectByName(name string) (ents.Foundation, error)
	SelectByCountry(country string) ([]ents.Foundation, error)
}

type FoundationService struct {
	FR IFoundationRepository //	FoundationRepository
}

func NewFoundationService(frepo IFoundationRepository) *FoundationService {
	return &FoundationService{FR: frepo}
}
func (FS *FoundationService) GetRepo() IFoundationRepository {
	return FS.FR
}
func (FS *FoundationService) ExistsById(id uint64) bool {
	_, result := FS.FR.SelectById(id)
	return (result == nil)
}
func (FS *FoundationService) ExistsByLogin(name string) bool {
	_, result := FS.FR.SelectByLogin(name)
	return (result == nil)
}
func (FS *FoundationService) ExistsByName(name string) bool {
	_, result := FS.FR.SelectByName(name)
	return (result == nil)
}

func (FS *FoundationService) ExistsByCountry(country string) bool {
	ans, err := FS.FR.SelectByCountry(country)
	return (err == nil && len(ans) != 0)
}

func (FS *FoundationService) Add(FPars ents.FoundationAdd) (ents.Foundation, error) {
	if FS.ExistsByLogin(FPars.Login) {
		return ents.Foundation{}, my_errors.ErrorConflict
	}
	var F ents.Foundation = ents.NewFoundation()
	var err_ error = validation.CheckFoundationAddParams(FPars)
	if err_ != nil {
		return ents.Foundation{}, err_
	}
	F.SetLogin(FPars.Login)
	F.SetPassword(FPars.Password)
	F.SetCountry(FPars.Country)
	F.SetName(FPars.Name)

	return FS.FR.Insert(F)
}

func (FS *FoundationService) Update(id_ string, FPars ents.FoundationAdd) (ents.Foundation, error) {
	if !FS.ExistsByLogin(FPars.Login) {
		return ents.Foundation{}, fmt.Errorf(my_errors.ErrNotExists)
	}
	var err error
	var F ents.Foundation
	F, err = FS.GetById(id_)
	if err != nil {
		return ents.Foundation{}, err
	}
	err = validation.CheckFoundationAddParams(FPars)
	if err != nil {
		return ents.Foundation{}, err
	}
	F.SetLogin(FPars.Login)
	F.SetPassword(FPars.Password)
	F.SetCountry(FPars.Country)
	F.SetName(FPars.Name)

	return FS.FR.Update(F)
}

func (FS *FoundationService) Delete(id_ string) (ents.Foundation, error) {
	F, err := FS.GetById(id_)
	if err != nil {
		return ents.Foundation{}, err
	}

	return FS.FR.Delete(F)
}
func (FS *FoundationService) GetAll() ([]ents.Foundation, error) {
	Foundations, err := FS.FR.Select()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список всех фондов: %s", err)
	}
	return Foundations, nil
}

func (FS *FoundationService) GetById(id_ string) (ents.Foundation, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return ents.Foundation{}, fmt.Errorf("некорректный id")
	}
	if !FS.ExistsById(id) {
		return ents.Foundation{}, fmt.Errorf(my_errors.ErrNotExists)
	}
	F, err := FS.FR.SelectById(id)
	if err != nil {
		return F, fmt.Errorf("не удалось получить фонд по id: %s", err)
	}
	return F, nil
}
func (FS *FoundationService) GetByLogin(name_ string) (ents.Foundation, error) {
	var F ents.Foundation
	if !FS.ExistsByLogin(name_) {
		return F, fmt.Errorf(my_errors.ErrNotExists)
	} else {
		var err error
		F, err = FS.FR.SelectByLogin(name_)
		if err != nil {
			return F, fmt.Errorf("не удалось получить фонд по логину: %s", err)
		}
	}
	return F, nil
}
func (FS *FoundationService) GetByName(name_ string) (ents.Foundation, error) {
	var F ents.Foundation
	if !FS.ExistsByName(name_) {
		return F, fmt.Errorf(my_errors.ErrNotExists)
	} else {
		var err error
		F, err = FS.FR.SelectByName(name_)
		if err != nil {
			return F, fmt.Errorf("не удалось получить фонд по названию: %s", err)
		}
	}
	return F, nil
}

func (FS *FoundationService) GetByCountry(country string) ([]ents.Foundation, error) {
	var F []ents.Foundation
	if validation.CheckCountry(country) != nil {
		return F, fmt.Errorf(my_errors.ErrCountry)
	}

	F, err := FS.FR.SelectByCountry(country)
	if err != nil {
		return F, fmt.Errorf("не удалось получить фонд по стране" + err.Error())
	}
	return F, nil
}

func (FS *FoundationService) Donate(F *ents.Foundation, sum float64) error {

	F.Fund_balance -= sum
	F.Outcome_history += sum

	_, err := FS.FR.Update(*F)
	return err
}

func (FS *FoundationService) AcceptDonate(id_ string, sum float64) error {

	var F ents.Foundation
	var err error
	F, err = FS.GetById(id_)
	if err != nil {
		return fmt.Errorf("фонду не удалось получить средства: %s", err)
	}
	F.Fund_balance += sum
	F.Income_history += sum
	_, err = FS.FR.Update(F)
	return err
}
func (FS *FoundationService) ReplenishBalance(U *ents.Foundation, sum float64) error {
	var err error
	if err != nil {
		return fmt.Errorf("фонду не удалось пополнить баланс: %s", err)
	}
	U.Fund_balance += sum
	_, err = FS.FR.Update(*U)
	return err
}
