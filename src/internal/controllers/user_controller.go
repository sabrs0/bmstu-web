package controllers

import (
	"fmt"
	"strconv"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/internal/business/validation"
)

type IUserService interface {
	Add(UPars ents.UserAdd) error
	Update(id_ string, UPars ents.UserAdd) error
	Delete(id_ string) error
	GetAll() ([]ents.User, error)
	GetById(id_ string) (ents.User, error)
	GetByLogin(id_ string) (ents.User, error)
	ExistsById(id uint64) bool
	ExistsByLogin(name string) bool
	Donate(U *ents.User, DP ents.UserDonate) error
	ReplenishBalance(U *ents.User, sum float64) error
	GetRepo() servs.IUserRepository
}
type UserController struct {
	TS  servs.TransactionService
	FS  servs.FoundationService
	FgS servs.FoundrisingService
	US  servs.UserService
}

func NewUserController(FdR servs.IFoundationRepository, FgR servs.IFoundrisingRepository,
	UR servs.IUserRepository, TR servs.ITransactionRepository) *UserController {
	return &UserController{
		TS:  *servs.NewTransactionService(TR),
		FS:  *servs.NewFoundationService(FdR),
		FgS: *servs.NewFoundrisingService(FgR),
		US:  *servs.NewUserService(UR),
	}
}
func (UC *UserController) GetAll() ([]ents.User, error) {
	return UC.US.GetAll()
}
func (UC *UserController) GetByID(id_ string) (ents.User, error) {
	return UC.US.GetById(id_)
}
func (UC *UserController) GetByLogin(login string) (ents.User, error) {
	return UC.US.GetByLogin(login)
}

func (UC *UserController) Add(params ents.UserAdd) (ents.User, error) {
	return UC.US.Add(params)

}

func (UC *UserController) Delete(id string) (ents.User, error) {
	u, err := UC.US.Delete(id)
	if err != nil {
		return ents.User{}, err
	}
	transactions1, err := UC.TS.GetFromId(ents.FROM_USER, id, &UC.FS, &UC.US)
	if err != nil {
		return ents.User{}, err
	}
	for i := range transactions1 {
		_, err = UC.TS.Delete(strconv.FormatUint(transactions1[i].Id, 10))
		if err != nil {
			return ents.User{}, err
		}
	}
	return u, err

}
func (UC *UserController) Update(id string, params ents.UserAdd) (ents.User, error) {
	var User ents.User
	User, _ = UC.GetByID(id)
	if params.Login == "" {
		params.Login = User.Login
	}
	if params.Password == "" {
		params.Password = User.Password
	}
	return UC.US.Update(id, params)

}
func (UC *UserController) Donate(userID string, params ents.UserDonate) (ents.Transaction, error) {
	switch params.EntityType {
	case ents.TO_FOUNDATION:
		return UC.DonateToFoundation(userID, params)
	case ents.TO_FOUNDRISING:
		return UC.DonateToFoundrising(userID, params)
	}
	return ents.Transaction{}, nil
}
func (UC *UserController) DonateToFoundation(userID string,
	params ents.UserDonate) (transaction ents.Transaction, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось выполнить пожертвование фонду от пользователя: %w", err)
		}

	}()

	err = validation.CheckMoneyFormat(params.SumOfMoney)
	if err != nil {
		return ents.Transaction{}, err
	}
	sum, err := strconv.ParseFloat(params.SumOfMoney, 64)
	if err != nil {
		return ents.Transaction{}, err
	}
	U, err := UC.GetByID(userID)
	if err != nil {
		return ents.Transaction{}, err
	}
	if U.Balance < sum {
		err = fmt.Errorf("недостаточно средств ")
		return ents.Transaction{}, err
	}
	err = UC.US.Donate(&U, sum)
	if err != nil {
		return ents.Transaction{}, err
	}
	err = UC.FS.AcceptDonate(params.EntityID, sum)
	if err != nil {
		return ents.Transaction{}, err
	}
	sid, err := strconv.Atoi(params.EntityID)
	found_id := uint64(sid)
	if err != nil {
		return ents.Transaction{}, err
	}
	TP := ents.TransactionAdd{
		From_essence_type: ents.FROM_USER,
		From_id:           U.Id,
		To_essence_type:   ents.TO_FOUNDATION,
		Sum_of_money:      sum,
		Comment:           params.Comm,
		To_id:             found_id,
	}
	transaction, err = UC.TS.Add(TP)
	if err != nil {
		return ents.Transaction{}, err
	}
	return transaction, nil
}
func isNewPhilantrop(TS servs.ITransactionService, FgS servs.FoundrisingService,
	foundrisingId string, userId uint64) (bool, error) {
	IDs, err := TS.GetFoundrisingPhilantropIds(foundrisingId, &FgS)
	var isNew bool
	if err != nil {
		return false, err
	}
	for _, id := range IDs {
		if id == userId {
			isNew = true
		}
	}
	return isNew, err
}

func (UC *UserController) DonateToFoundrising(userID string,
	params ents.UserDonate) (transaction ents.Transaction, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось выполнить пожертвование сбору от пользователя: %w", err)
		}

	}()

	err = validation.CheckMoneyFormat(params.SumOfMoney)

	if err != nil {
		return ents.Transaction{}, err
	}
	var reqSum float64
	reqSum, err = strconv.ParseFloat(params.SumOfMoney, 64)
	if err != nil {
		return ents.Transaction{}, err
	}
	U, err := UC.GetByID(userID)
	if err != nil {
		return ents.Transaction{}, err
	}
	if U.Balance < reqSum {
		return ents.Transaction{}, fmt.Errorf("недостаточно средств ")
	}
	foundrising, err := UC.FgS.GetById(params.EntityID)
	if err != nil {
		return ents.Transaction{}, err
	}
	if !foundrising.Closing_date.Valid {
		err = UC.US.Donate(&U, reqSum)
		if err != nil {
			return ents.Transaction{}, err
		}
		isNewPh, err := isNewPhilantrop(&UC.TS, UC.FgS, params.EntityID, U.Id)
		if err != nil {
			return ents.Transaction{}, err
		}
		var remainder float64
		remainder, err = UC.FgS.AcceptDonate(params.EntityID, reqSum, isNewPh)
		if err != nil {
			return ents.Transaction{}, err
		}
		sid, err := strconv.Atoi(params.EntityID)
		foundrising_id := uint64(sid)
		if err != nil {
			return ents.Transaction{}, err
		}
		TP := ents.TransactionAdd{
			From_essence_type: ents.FROM_USER,
			From_id:           U.Id,
			To_essence_type:   ents.TO_FOUNDRISING,
			Sum_of_money:      reqSum,
			Comment:           params.Comm,
			To_id:             foundrising_id,
		}
		transaction, err = UC.TS.Add(TP)
		if err != nil {
			return ents.Transaction{}, err
		}
		if remainder > 0.0 {
			foundrising, _ := UC.FgS.GetById(params.EntityID)
			found_id := foundrising.Found_id
			TP := ents.TransactionAdd{
				From_essence_type: ents.FROM_USER,
				From_id:           U.Id,
				To_essence_type:   ents.TO_FOUNDATION,
				Sum_of_money:      remainder,
				Comment:           "returning the remain",
				To_id:             found_id,
			}
			_, err = UC.TS.Add(TP)
			return ents.Transaction{}, err
		} else if remainder <= 1e-9 {
			U.ClosedFingAmount += 1
			_, err = UC.US.GetRepo().Update(U)
			return ents.Transaction{}, err
		}
	}
	return transaction, nil
}

func (UC *UserController) ReplenishBalance(sum string, U *ents.User) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось пополнить баланс:  %w", err)
		}

	}()
	err = validation.CheckMoneyFormat(sum)
	if err != nil {
		return err
	}
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	if err != nil {
		return err
	}

	if reqSum > 50000.00 {
		return fmt.Errorf("введенная сумма превышается 50 000")
	}
	return UC.US.ReplenishBalance(U, reqSum)
}
