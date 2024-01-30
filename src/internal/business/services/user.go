package services

import (
	"fmt"
	"strconv"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	"github.com/sabrs0/bmstu-web/src/internal/business/validation"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IUserRepository interface {
	Insert(U ents.User) (ents.User, error)
	Update(U ents.User) (ents.User, error)
	Delete(U ents.User) (ents.User, error)
	Select() ([]ents.User, error)
	SelectById(id uint64) (ents.User, error)
	SelectByLogin(name string) (ents.User, error)
	//GetDB() *gorm.DB
}

type UserService struct {
	UR IUserRepository
}

func NewUserService(urepo IUserRepository) *UserService {
	return &UserService{UR: urepo}
}
func (US *UserService) GetRepo() IUserRepository {
	return US.UR
}
func (US *UserService) ExistsById(id uint64) bool {
	_, result := US.UR.SelectById(id)
	return (result == nil)
}

func (US *UserService) ExistsByLogin(name string) bool {
	_, result := US.UR.SelectByLogin(name)
	return (result == nil)
}

func (US *UserService) Add(UPars ents.UserAdd) (ents.User, error) {
	if US.ExistsByLogin(UPars.Login) {
		return ents.User{}, my_errors.ErrorConflict
	}
	U := ents.NewUser()
	err := validation.CheckUserAddParams(UPars)
	if err != nil {
		return ents.User{}, err
	}
	U.SetLogin(UPars.Login)
	U.SetPassword(UPars.Password)

	return US.UR.Insert(U)
}

func (US *UserService) Update(id_ string, UPars ents.UserAdd) (ents.User, error) {

	var errGet error
	var U ents.User
	U, errGet = US.GetById(id_)
	if errGet != nil {
		return ents.User{}, errGet
	}
	if US.ExistsByLogin(UPars.Login) && U.Login != UPars.Login {
		return ents.User{}, my_errors.ErrorConflict
	}
	err := validation.CheckUserAddParams(UPars)
	if err != nil {
		return ents.User{}, err
	}
	U.SetLogin(UPars.Login)
	U.SetPassword(UPars.Password)

	return US.UR.Update(U)
}

func (US *UserService) Delete(id_ string) (ents.User, error) {
	var errGet error
	var U ents.User
	U, errGet = US.GetById(id_)
	if errGet != nil {
		return ents.User{}, errGet
	}
	return US.UR.Delete(U)
}
func (US *UserService) GetAll() ([]ents.User, error) {
	Users, err := US.UR.Select()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список всех пользователей: %s", err)
	}
	return Users, nil
}
func (US *UserService) GetById(id_ string) (ents.User, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return ents.User{}, fmt.Errorf("некорректный id=%s, %s", id_, err.Error())
	}
	if !US.ExistsById(id) {
		return ents.User{}, fmt.Errorf(my_errors.ErrNotExists)
	}
	U, err := US.UR.SelectById(id)
	if err != nil {
		return U, fmt.Errorf("не удалось получить пользователся по id:  %w", my_errors.ErrorNotExists)
	}
	return U, nil
}

func (US *UserService) GetByLogin(name_ string) (ents.User, error) {

	if !US.ExistsByLogin(name_) {
		return ents.User{}, fmt.Errorf(my_errors.ErrNotExists)
	}
	U, err := US.UR.SelectByLogin(name_)
	if err != nil {
		return U, fmt.Errorf("не удалось получить пользователся по логину: %w", my_errors.ErrorNotExists)
	}
	return U, nil
}

// проверка формата денег и логики самих денег будет на уровне controllers, сюда приходят уже проверенные деньги
func (US *UserService) Donate(U *ents.User, sum float64) error {

	U.Balance -= sum
	U.CharitySum += sum
	/*if DP.IsClosedFoundrising {
		U.ClosedFingAmount += 1
	}*/
	_, err := US.UR.Update(*U)
	return err
}

func (US *UserService) ReplenishBalance(U *ents.User, sum float64) (ents.UserTransfer, error) {
	U.Balance += sum

	_, err := US.UR.Update(*U)
	return ents.NewUserTransfer(*U), err
}
