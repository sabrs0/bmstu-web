package businessLogicTest

import (
	"math"
	"reflect"
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	chk "github.com/sabrs0/bmstu-web/src/internal/business/validation"
)

func TestUserServiceAdd(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	pars := chk.NewUserMainParams("user2", "1234")
	err := serv.Add(pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestUserServiceUpdate(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	pars := chk.NewUserMainParams("user3", "123")
	err := serv.Update("1", pars)
	if err == nil {
		t.Errorf("Shoud be error, but err is nil\n")
	}
	pars = chk.NewUserMainParams("user", "123")
	err = serv.Update("0", pars)
	if err == nil {
		t.Errorf("Shoud be error, but err is nil\n")
	}
	pars = chk.NewUserMainParams("user2", "123")
	err = serv.Update("0", pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestUserServiceGetAll(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error at getAllUsers: %s\n", err)
	}
	if !reflect.DeepEqual(Users, ans) {
		t.Errorf("Shoud be equal getAll. ans = %v, src = %v\n", ans, Users)
	}
}
func TestUserServiceGetById(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error at getById: %s\n", err)
	}
	if ans != Users[0] {
		t.Errorf("Should be equal getById. ans = %v, src = %v\n", ans, Users)
	}
}
func GetByLogin(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	ans, err := serv.GetByLogin("user")
	if err != nil {
		t.Errorf("Some error at getByLogin: %s\n", err)
	}
	if !reflect.DeepEqual(Users, ans) {
		t.Errorf("Should be equal getByLogin. ans = %v, src = %v\n", ans, Users)
	}
}
func TestUserServiceDonate(t *testing.T) {
	var Users []ents.User = []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    110,
			CharitySum: 1000,
		},
	}
	serv := servs.NewUserService(NewUserRepMock(Users))
	pars := chk.NewUserDonateParams(100.00, false)
	err := serv.Donate(&Users[0], pars)
	if err != nil {
		t.Errorf("Some error at Donate: %s\n", err)
	}
	if math.Abs(Users[0].Balance-10.00) > 1e-9 {
		t.Errorf("Incorrect balance after donate: balance - %f\n", Users[0].Balance)
	}
}
