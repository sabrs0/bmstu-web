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
func (UC *UserController) GetAll() ([]ents.UserTransfer, error) {
	users, err := UC.US.GetAll()
	if err != nil {
		return nil, err
	}
	uTrans := []ents.UserTransfer{}
	for _, u := range users {
		uTrans = append(uTrans, ents.NewUserTransfer(u))
	}
	return uTrans, nil
}
func (UC *UserController) GetByID(id_ string) (ents.UserTransfer, error) {
	user, err := UC.US.GetById(id_)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	return ents.NewUserTransfer(user), nil
}
func (UC *UserController) GetByLogin(login string) (ents.UserTransfer, error) {
	user, err := UC.US.GetByLogin(login)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	return ents.NewUserTransfer(user), nil

}

func (UC *UserController) Add(params ents.UserAdd) (ents.UserTransfer, error) {
	user, err := UC.US.Add(params)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	return ents.NewUserTransfer(user), nil
}

func (UC *UserController) Delete(id string) (ents.UserTransfer, error) {
	u, err := UC.US.Delete(id)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	transactions1, err := UC.TS.GetFromId(ents.FROM_USER, id, &UC.FS, &UC.US)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	for i := range transactions1 {
		_, err = UC.TS.Delete(strconv.FormatUint(transactions1[i].Id, 10))
		if err != nil {
			return ents.UserTransfer{}, err
		}
	}
	return ents.NewUserTransfer(u), nil

}
func (UC *UserController) Update(id string, params ents.UserAdd) (ents.UserTransfer, error) {
	var User ents.User
	User, _ = UC.US.GetById(id)
	if params.Login == "" {
		params.Login = User.Login
	}
	if params.Password == "" {
		params.Password = User.Password
	}
	user, err := UC.US.Update(id, params)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	return ents.NewUserTransfer(user), nil

}
func (UC *UserController) Donate(userID string, params ents.UserDonate) (ents.TransactionTransfer, error) {
	switch params.EntityType {
	case ents.TO_FOUNDATION:

		tr, err := UC.DonateToFoundation(userID, params)
		if err != nil {
			return ents.TransactionTransfer{}, err
		}
		return tr, nil
	case ents.TO_FOUNDRISING:
		tr, err := UC.DonateToFoundrising(userID, params)
		if err != nil {
			return ents.TransactionTransfer{}, err
		}
		return tr, nil
	}
	return ents.TransactionTransfer{}, nil
}
func (UC *UserController) DonateToFoundation(userID string,
	params ents.UserDonate) (transactionTr ents.TransactionTransfer, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось выполнить пожертвование фонду от пользователя: %w", err)
		}

	}()

	err = validation.CheckMoneyFormat(params.SumOfMoney)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	sum, err := strconv.ParseFloat(params.SumOfMoney, 64)
	if err != nil || params.SumOfMoney == "NaN" || params.SumOfMoney == "Infinity" || params.SumOfMoney == "Inf" {
		return ents.TransactionTransfer{}, err
	}
	U, err := UC.US.GetById(userID)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	if U.Balance < sum {
		err = fmt.Errorf("недостаточно средств ")
		return ents.TransactionTransfer{}, err
	}
	err = UC.US.Donate(&U, sum)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	err = UC.FS.AcceptDonate(params.EntityID, sum)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	sid, err := strconv.Atoi(params.EntityID)
	found_id := uint64(sid)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	TP := ents.TransactionAdd{
		From_essence_type: ents.FROM_USER,
		From_id:           U.Id,
		To_essence_type:   ents.TO_FOUNDATION,
		Sum_of_money:      sum,
		Comment:           params.Comm,
		To_id:             found_id,
	}
	transaction, err := UC.TS.Add(TP)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	return ents.NewTransactionTransfer(transaction), nil
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
	params ents.UserDonate) (transactionTr ents.TransactionTransfer, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось выполнить пожертвование сбору от пользователя: %w", err)
		}

	}()

	err = validation.CheckMoneyFormat(params.SumOfMoney)

	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	var reqSum float64
	reqSum, err = strconv.ParseFloat(params.SumOfMoney, 64)
	if err != nil || params.SumOfMoney == "NaN" || params.SumOfMoney == "Infinity" || params.SumOfMoney == "Inf" {
		return ents.TransactionTransfer{}, err
	}
	U, err := UC.US.GetById(userID)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	if U.Balance < reqSum {
		return ents.TransactionTransfer{}, fmt.Errorf("недостаточно средств ")
	}
	foundrising, err := UC.FgS.GetById(params.EntityID)
	if err != nil {
		return ents.TransactionTransfer{}, err
	}
	var transaction ents.Transaction
	if !foundrising.Closing_date.Valid {
		err = UC.US.Donate(&U, reqSum)
		if err != nil {
			return ents.TransactionTransfer{}, err
		}
		isNewPh, err := isNewPhilantrop(&UC.TS, UC.FgS, params.EntityID, U.Id)
		if err != nil {
			return ents.TransactionTransfer{}, err
		}
		var remainder float64
		remainder, err = UC.FgS.AcceptDonate(params.EntityID, reqSum, isNewPh)
		if err != nil {
			return ents.TransactionTransfer{}, err
		}
		sid, err := strconv.Atoi(params.EntityID)
		foundrising_id := uint64(sid)
		if err != nil {
			return ents.TransactionTransfer{}, err
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
			return ents.TransactionTransfer{}, err
		}
		if remainder > 1e-9 {
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
			return ents.TransactionTransfer{}, err
		} else if remainder <= 1e-9 && remainder > -1.0 {
			U.ClosedFingAmount += 1
			_, err = UC.US.GetRepo().Update(U)
			return ents.TransactionTransfer{}, err
		}
	} else {
		return ents.TransactionTransfer{}, fmt.Errorf("данный сбор уже закрыт")
	}
	return ents.NewTransactionTransfer(transaction), nil
}

func (UC *UserController) ReplenishBalance(id string, params ents.UserReplenish) (userTr ents.UserTransfer, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Не удалось пополнить баланс:  %w", err)
		}

	}()
	var U ents.User
	U, err = UC.US.GetById(id)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	err = validation.CheckMoneyFormat(params.Sum_of_Money)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	var reqSum float64
	reqSum, err = strconv.ParseFloat(params.Sum_of_Money, 64)
	if err != nil || params.Sum_of_Money == "NaN" || params.Sum_of_Money == "Infinity" || params.Sum_of_Money == "Inf" {
		return ents.UserTransfer{}, err
	}

	if reqSum > 50000.00 {
		return ents.UserTransfer{}, fmt.Errorf("введенная сумма превышается 50 000")
	}
	_, err = UC.US.ReplenishBalance(&U, reqSum)
	if err != nil {
		return ents.UserTransfer{}, err
	}
	return ents.NewUserTransfer(U), nil
}
