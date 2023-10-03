package businessLogicTest

import (
	"math"
	"reflect"
	"testing"

	chk "github.com/sabrs0/bmstu-web/src/internal/business/checker"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
)

func TestFoundationServiceAdd(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:       1,
			Name:     "Foundation",
			Password: "123",
			Country:  "Russia",
			Login:    "Foundation123",
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	pars := chk.NewFoundationMainParams("Foundation123", "123", "F", "Russia")
	err := serv.Add(pars)
	if err == nil {
		t.Errorf("Shoud be error, but err is nil\n")
	}
	pars = chk.NewFoundationMainParams("Foundation124", "123", "F", "Russia")
	err = serv.Add(pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestFoundationServiceUpdate(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:       1,
			Name:     "Foundation",
			Password: "123",
			Country:  "Russia",
			Login:    "Foundation123",
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	pars := chk.NewFoundationMainParams("Foundation123", "123", "F", "Russia")
	err := serv.Update("0", pars)
	if err == nil {
		t.Errorf("Shoud be error, but err is nil\n")
	}
	pars = chk.NewFoundationMainParams("Foundation125", "123", "F", "Russia")
	err = serv.Update("1", pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestFoundationServiceGetAll(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:       1,
			Name:     "Foundation",
			Password: "123",
			Country:  "Russia",
			Login:    "Foundation123",
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error at getAllFoudnations: %s\n", err)
	}
	if !reflect.DeepEqual(Foundations, ans) {
		t.Errorf("Shoud be equal getAll. ans = %v, src = %v\n", ans, Foundations)
	}
}
func TestFoundationServiceGetById(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:       1,
			Name:     "Foundation",
			Password: "123",
			Country:  "Russia",
			Login:    "Foundation123",
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error at getAllFoudnations: %s\n", err)
	}
	if ans != Foundations[0] {
		t.Errorf("Should be equal getById. ans = %v, src = %v\n", ans, Foundations)
	}
}
func TestFoundationServiceGetBylogin(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:       1,
			Name:     "Foundation",
			Password: "123",
			Country:  "Russia",
			Login:    "Foundation123",
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	ans, err := serv.GetByLogin("Foundation123")
	if err != nil {
		t.Errorf("Some error at getAllFoudnations: %s\n", err)
	}
	if ans != Foundations[0] {
		t.Errorf("Should be equal getById. ans = %v, src = %v\n", ans, Foundations)
	}
	_, err = serv.GetByLogin("Foundation124")
	if err == nil {
		t.Errorf("Should be nil at getAllFoudnations:\n")
	}
}
func TestFoundationServiceDonate(t *testing.T) {
	var Foundations []ents.Foundation = []ents.Foundation{
		{
			Id:           1,
			Name:         "Foundation",
			Password:     "123",
			Country:      "Russia",
			Login:        "Foundation123",
			Fund_balance: 110.00,
		},
	}
	serv := servs.NewFoundationService(NewFoundationRepMock(Foundations))
	pars := chk.NewFoundationDonateParams(100.00, false)
	err := serv.Donate(&Foundations[0], pars)
	if err != nil {
		t.Errorf("Some error at Donate: %s\n", err)
	}
	if math.Abs(Foundations[0].Fund_balance-10.00) > 1e-9 {
		t.Errorf("Incorrect balance after donate: balance - %f\n", Foundations[0].Fund_balance)
	}
}
