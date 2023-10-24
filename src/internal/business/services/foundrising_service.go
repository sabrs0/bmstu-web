package services

import (
	"fmt"
	"math"
	"strconv"
	"time"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type IFoundrisingRepository interface {
	Insert(F ents.Foundrising) (ents.Foundrising, error)
	Update(F ents.Foundrising) (ents.Foundrising, error)
	Delete(F ents.Foundrising) (ents.Foundrising, error)
	Select() ([]ents.Foundrising, error)
	SelectById(id uint64) (ents.Foundrising, error)
	SelectByFoundId(id uint64) ([]ents.Foundrising, error)
	SelectByCreateDate(date string) ([]ents.Foundrising, error)
	SelectByFoundIdActive(id uint64) ([]ents.Foundrising, error)
	SelectByCloseDate(date string) ([]ents.Foundrising, error)
	SelectByIdAndFoundId(id_ uint64, found_id_ uint64) (ents.Foundrising, error)
	//GetDB() *gorm.DB
}

type FoundrisingService struct {
	FR IFoundrisingRepository
}

func NewFoundrisingService(frepo IFoundrisingRepository) *FoundrisingService {
	return &FoundrisingService{FR: frepo}
}
func (FS *FoundrisingService) GetRepo() IFoundrisingRepository {
	return FS.FR
}

func (FS *FoundrisingService) ExistsById(id uint64) bool {
	_, result := FS.FR.SelectById(id)
	return (result == nil)
}

func (FS *FoundrisingService) Add(FPars ents.FoundrisingAdd) (ents.Foundrising, error) {

	var U ents.Foundrising = ents.NewFoundrising()
	floatSum, _ := strconv.ParseFloat(FPars.Required_sum, 64)
	U.SetReqSum(floatSum)
	U.SetCreateDate(FPars.Creation_date)
	U.SetDescription(FPars.Description)
	U.SetFoundId(FPars.Found_id)

	return FS.FR.Insert(U)
}

func (FS *FoundrisingService) Update(id_ string, FPars ents.FoundrisingPut) (ents.Foundrising, error) {

	U, err := FS.GetById(id_)
	if err != nil {
		return ents.Foundrising{}, err
	}
	U.SetDescription(FPars.Description)
	floatSum, _ := strconv.ParseFloat(FPars.Required_sum, 64)

	U.SetReqSum(floatSum)

	return FS.FR.Update(U)
}

func (FS *FoundrisingService) Delete(id_ string) (ents.Foundrising, error) {
	U, err := FS.GetById(id_)
	if err != nil {
		return ents.Foundrising{}, err
	}
	return FS.FR.Delete(U)
}
func (FS *FoundrisingService) GetAll() ([]ents.Foundrising, error) {
	Foundrisings, err := FS.FR.Select()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список всех сборов: %s", err)
	} else {
		return Foundrisings, nil
	}
}
func (FS *FoundrisingService) GetById(id_ string) (ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U ents.Foundrising
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	}
	if !FS.ExistsById(id) {
		return U, fmt.Errorf("несуществующий id %w", my_errors.ErrorNotExists)
	}
	U, err = FS.FR.SelectById(id)
	if err != nil {
		return U, fmt.Errorf("не удалось получить сбор по id: %s", err)
	}
	return U, nil
}

func (FS *FoundrisingService) GetByFoundId(id_ string) ([]ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return nil, fmt.Errorf("некорректный id %w", my_errors.ErrorNotExists)
	}
	U, err := FS.FR.SelectByFoundId(id)
	if err != nil {
		return U, fmt.Errorf("не удалось получить сбор по id фонда: %s", err)
	}
	return U, nil
}

func (FS *FoundrisingService) GetByIdAndFoundId(id_ string, found_id_ string) (ents.Foundrising, error) {
	sid, err1 := strconv.Atoi(id_)
	id := uint64(sid)

	s_found_id, err := strconv.Atoi(found_id_)

	var U ents.Foundrising
	if err != nil || err1 != nil {
		return U, fmt.Errorf("некорректный id")
	}
	found_id := uint64(s_found_id)
	U, err = FS.FR.SelectByIdAndFoundId(id, found_id)
	if err != nil {
		return U, fmt.Errorf("не удалось получить фонд данного сбора: %s", err)
	}
	return U, nil
}

func (FS *FoundrisingService) GetByFoundIdActive(id_ string) ([]ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	if err != nil {
		return nil, fmt.Errorf("некорректный id")
	}
	U, err := FS.FR.SelectByFoundIdActive(id)
	if err != nil {
		return U, fmt.Errorf(my_errors.ErrNotExistsFoundrising)
	}

	return U, nil
}

func (FS *FoundrisingService) GetByCreateDate(date string) ([]ents.Foundrising, error) {
	U, err := FS.FR.SelectByCreateDate(date)
	if err != nil {
		return U, fmt.Errorf("не удалось получить сбор по id по дате создания: %s", err)
	}
	return U, nil
}

func (FS *FoundrisingService) GetByCloseDate(date string) ([]ents.Foundrising, error) {
	U, err := FS.FR.SelectByCloseDate(date)
	if err != nil {
		return U, fmt.Errorf("не удалось получить сбор по дате закрытия: %s", err)
	}
	return U, nil
}

func (FS *FoundrisingService) AcceptDonate(id_ string, sum float64, isNewPhil bool) (float64, error) {

	var remainder float64 = -1.0
	var F ents.Foundrising
	var err error
	F, err = FS.GetById(id_)
	if err != nil {
		return remainder, err
	}
	if isNewPhil {
		F.Philantrops_amount += 1
	}
	if F.Current_sum+sum > F.Required_sum {
		remainder = F.Current_sum + sum - F.Required_sum
		F.Closing_date.String = time.Now().Format(ents.DateFormat)
		F.Closing_date.Valid = true
		F.Current_sum = F.Required_sum
	} else if math.Abs(F.Current_sum+sum-F.Required_sum) <= 1e-9 {
		F.Closing_date.String = time.Now().Format(ents.DateFormat)
		F.Closing_date.Valid = true
		F.Current_sum = F.Required_sum
		remainder = 0.00
	} else {
		F.Current_sum += sum
	}
	_, err = FS.FR.Update(F)
	if err != nil {
		return 0.0, err
	}
	return remainder, nil
}
