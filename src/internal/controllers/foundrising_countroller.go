package controllers

import (
	"fmt"
	"strconv"
	"time"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/internal/business/validation"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type FoundrisingController struct {
	FS   servs.FoundrisingService
	FndS servs.FoundationService
}

func NewFoundrisingController(FgR servs.IFoundrisingRepository, FndR servs.IFoundationRepository) *FoundrisingController {
	return &FoundrisingController{
		FS:   *servs.NewFoundrisingService(FgR),
		FndS: *servs.NewFoundationService(FndR),
	}
}

func (UC *FoundrisingController) GetAll() ([]ents.Foundrising, error) {
	return UC.FS.GetAll()
}
func (UC *FoundrisingController) GetByID(id_ string) (ents.Foundrising, error) {
	return UC.FS.GetById(id_)
}
func (UC *FoundrisingController) GetByCreDate(date string) ([]ents.Foundrising, error) {
	return UC.FS.GetByCreateDate(date)
}
func (UC *FoundrisingController) GetByCloDate(date string) ([]ents.Foundrising, error) {
	return UC.FS.GetByCloseDate(date)
}
func (UC *FoundrisingController) GetByFoundId(id string) ([]ents.Foundrising, error) {
	return UC.FS.GetByFoundId(id)
}
func (UC *FoundrisingController) GetByFoundIdActive(id string) ([]ents.Foundrising, error) {
	return UC.FS.GetByFoundIdActive(id)
}
func (UC *FoundrisingController) GetByIdAndFoundId(id string, found_id string) (ents.Foundrising, error) {
	return UC.FS.GetByIdAndFoundId(id, found_id)
}

func (UC *FoundrisingController) Add(params ents.FoundrisingAdd) (ents.Foundrising, error) {
	//А может все таки в params found_id - string ?
	/*sid, err := strconv.Atoi(params.Found_id)
	found_id := uint64(sid)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf("некорректный id фонда")
	}*/
	err := validation.CheckMoneyFormat(params.Required_sum)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf(my_errors.ErrMoney)
	}
	_, err = strconv.ParseFloat(params.Required_sum, 64)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf(my_errors.ErrMoney)
	}
	create_date := time.Now().Format(ents.DateFormat)
	params.Creation_date = create_date
	if !UC.FndS.ExistsById(params.Found_id) {
		return ents.Foundrising{}, fmt.Errorf("фонда с таким ID не существует %w", my_errors.ErrorNotExists)
	}
	return UC.FS.Add(params)

}

func (UC *FoundrisingController) Delete(id string) (ents.Foundrising, error) {
	return UC.FS.Delete(id)

}
func (UC *FoundrisingController) Update(id string, params ents.FoundrisingPut) (ents.Foundrising, error) {
	var Foundrising ents.Foundrising
	Foundrising, _ = UC.GetByID(id)
	if params.Description == "" {
		params.Description = Foundrising.Description
	}
	if params.Required_sum == "" {
		params.Required_sum = strconv.FormatFloat(Foundrising.Required_sum, 'f', 2, 64)

	}
	err := validation.CheckMoneyFormat(params.Required_sum)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf(my_errors.ErrMoney)
	}
	reqSumfloat, err := strconv.ParseFloat(params.Required_sum, 64)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf(my_errors.ErrMoney)
	}

	if reqSumfloat < Foundrising.Required_sum {
		return ents.Foundrising{}, fmt.Errorf("новая сумма меньше той, что была прежде")
	}

	return UC.FS.Update(id, params)

}
