package businessLogicTest

import (
	"database/sql"
	"math"
	"reflect"
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	chk "github.com/sabrs0/bmstu-web/src/internal/business/checker"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
)

func TestFoundrisingServiceAdd(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	pars := chk.NewFoundrisingMainParams(1, "123", 100, "01-01-2002")
	err := serv.Add(pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestFoundrisingServiceUpdate(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	pars := chk.NewFoundrisingMainParams(1, "123", 100, "01-01-2001")
	err := serv.Update("0", pars)
	if err == nil {
		t.Errorf("Shoud be error, but err is nil\n")
	}
	pars = chk.NewFoundrisingMainParams(1, "123", 100, "01-01-2001")
	err = serv.Update("1", pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestFoundrisingServiceGetAll(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error at getAllFoundrisings: %s\n", err)
	}
	if !reflect.DeepEqual(Foundrisings, ans) {
		t.Errorf("Shoud be equal getAll. ans = %v, src = %v\n", ans, Foundrisings)
	}
}
func TestFoundrisingServiceGetById(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error at getById: %s\n", err)
	}
	if ans != Foundrisings[0] {
		t.Errorf("Should be equal getById. ans = %v, src = %v\n", ans, Foundrisings)
	}
}
func TestFoundrisingServiceGetByCreateDate(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	ans, err := serv.GetByCreateDate("01-01-2021")
	if err != nil {
		t.Errorf("Some error at getByCreateDate: %s\n", err)
	}
	if !reflect.DeepEqual(Foundrisings, ans) {
		t.Errorf("Should be equal getByCreateDate. ans = %v, src = %v\n", ans, Foundrisings)
	}
	ans, err = serv.GetByCreateDate("01-01-2022")
	if len(ans) != 0 {
		t.Errorf("Should be zero len at getByCreateDate:\n")
	}
}
func TestFoundrisingServiceAcceptDonate(t *testing.T) {
	var Foundrisings []ents.Foundrising = []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	serv := servs.NewFoundrisingService(NewFoundrisingRepMock(Foundrisings))
	rem, err := serv.AcceptDonate("1", 50, true)
	if err != nil {
		t.Errorf("Some error at Donate: %s\n", err)
	}
	if math.Abs(rem-(100-50)) > 1e-9 {
		t.Errorf("Incorrect remainder after acceptDonate: balance - %f\n", rem)
	}
}
