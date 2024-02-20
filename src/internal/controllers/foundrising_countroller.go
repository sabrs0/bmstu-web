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

func (UC *FoundrisingController) GetAll() ([]ents.FoundrisingTransfer, error) {
	fRisings, err := UC.FS.GetAll()
	if err != nil {
		return nil, err
	}
	fTransfer := []ents.FoundrisingTransfer{}
	for _, f := range fRisings {
		fTransfer = append(fTransfer, ents.NewFoundrisingTransfer(f))
	}
	return fTransfer, nil

}
func (UC *FoundrisingController) GetByID(id_ string) (ents.FoundrisingTransfer, error) {
	fRising, err := UC.FS.GetById(id_)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	return ents.FoundrisingTransfer(fRising), nil
}
func (UC *FoundrisingController) GetByCreDate(date string) ([]ents.FoundrisingTransfer, error) {
	fRisings, err := UC.FS.GetByCreateDate(date)
	if err != nil {
		return nil, err
	}
	fTransfer := []ents.FoundrisingTransfer{}
	for _, f := range fRisings {
		fTransfer = append(fTransfer, ents.NewFoundrisingTransfer(f))
	}
	return fTransfer, nil
}
func (UC *FoundrisingController) GetByCloDate(date string) ([]ents.FoundrisingTransfer, error) {

	fRisings, err := UC.FS.GetByCloseDate(date)
	if err != nil {
		return nil, err
	}
	fTransfer := []ents.FoundrisingTransfer{}
	for _, f := range fRisings {
		fTransfer = append(fTransfer, ents.NewFoundrisingTransfer(f))
	}
	return fTransfer, nil
}
func (UC *FoundrisingController) GetByFoundId(id string) ([]ents.FoundrisingTransfer, error) {
	fRisings, err := UC.FS.GetByFoundId(id)
	if err != nil {
		return nil, err
	}
	fTransfer := []ents.FoundrisingTransfer{}
	for _, f := range fRisings {
		fTransfer = append(fTransfer, ents.NewFoundrisingTransfer(f))
	}
	return fTransfer, nil
}
func (UC *FoundrisingController) GetByFoundIdActive(id string) ([]ents.FoundrisingTransfer, error) {
	fRisings, err := UC.FS.GetByFoundIdActive(id)
	if err != nil {
		return nil, err
	}
	fTransfer := []ents.FoundrisingTransfer{}
	for _, f := range fRisings {
		fTransfer = append(fTransfer, ents.NewFoundrisingTransfer(f))
	}
	return fTransfer, nil
}
func (UC *FoundrisingController) GetByIdAndFoundId(id string, found_id string) (ents.FoundrisingTransfer, error) {
	fRising, err := UC.FS.GetByIdAndFoundId(id, found_id)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	return ents.FoundrisingTransfer(fRising), nil
}

func (UC *FoundrisingController) Add(params ents.FoundrisingAdd) (ents.FoundrisingTransfer, error) {
	err := validation.CheckMoneyFormat(params.Required_sum)
	if err != nil {
		return ents.FoundrisingTransfer{}, fmt.Errorf(my_errors.ErrMoney)
	}
	_, err = strconv.ParseFloat(params.Required_sum, 64)
	if err != nil || params.Required_sum == "NaN" || params.Required_sum == "Infinity" || params.Required_sum == "Inf" {
		return ents.FoundrisingTransfer{}, fmt.Errorf(my_errors.ErrMoney)
	}
	create_date := time.Now().Format(ents.DateFormat)
	params.Creation_date = create_date
	if !UC.FndS.ExistsById(params.Found_id) {
		return ents.FoundrisingTransfer{}, fmt.Errorf("фонда с таким ID не существует %w", my_errors.ErrorNotExists)
	}
	fRising, err := UC.FS.Add(params)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	strID := strconv.FormatUint(params.Found_id, 10)
	foundation, err := UC.FndS.GetById(strID)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	foundation.CurFoudrisingAmount += 1
	_, err = UC.FndS.FR.Update(foundation)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}

	return ents.FoundrisingTransfer(fRising), nil
}

func (UC *FoundrisingController) Delete(id string) (ents.FoundrisingTransfer, error) {
	fRising, err := UC.FS.Delete(id)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	// удалили незакрытый сбор - надо подчистить у foundation cur fing amount
	if !fRising.Closing_date.Valid || len(fRising.Closing_date.String) == 0 {
		strID := strconv.FormatUint(fRising.Found_id, 10)
		foundation, err := UC.FndS.GetById(strID)
		if err != nil {
			return ents.FoundrisingTransfer{}, err
		}
		if foundation.CurFoudrisingAmount > 0 {
			foundation.CurFoudrisingAmount -= 1
		}

		_, err = UC.FndS.FR.Update(foundation)
		if err != nil {
			return ents.FoundrisingTransfer{}, err
		}
	}
	return ents.FoundrisingTransfer(fRising), nil

}
func (UC *FoundrisingController) Update(id string, params ents.FoundrisingPut) (ents.FoundrisingTransfer, error) {
	var Foundrising ents.Foundrising
	Foundrising, _ = UC.FS.GetById(id)
	if params.Description == "" {
		params.Description = Foundrising.Description
	}
	if params.Required_sum == "" {
		params.Required_sum = strconv.FormatFloat(Foundrising.Required_sum, 'f', 2, 64)

	}
	err := validation.CheckMoneyFormat(params.Required_sum)
	if err != nil {
		return ents.FoundrisingTransfer{}, fmt.Errorf(my_errors.ErrMoney)
	}
	reqSumfloat, err := strconv.ParseFloat(params.Required_sum, 64)
	if err != nil || params.Required_sum == "NaN" || params.Required_sum == "Infinity" || params.Required_sum == "Inf" {
		return ents.FoundrisingTransfer{}, fmt.Errorf(my_errors.ErrMoney)
	}

	if reqSumfloat < Foundrising.Required_sum {
		return ents.FoundrisingTransfer{}, fmt.Errorf("новая сумма меньше той, что была прежде")
	}
	fRising, err := UC.FS.Update(id, params)
	if err != nil {
		return ents.FoundrisingTransfer{}, err
	}
	return ents.FoundrisingTransfer(fRising), nil

}
