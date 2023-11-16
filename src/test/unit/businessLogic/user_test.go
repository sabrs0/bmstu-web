package businessLogicTest

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
	"github.com/sabrs0/bmstu-web/src/test/unit/businessLogic/mocks"
)

func TestUserServiceAdd(t *testing.T) {
	expectedValue := ents.User{
		Id:       0,
		Login:    "user",
		Password: "123",
	}
	params := ents.UserAdd{
		Login:    "user",
		Password: "123",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().Insert(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("user").Return(ents.User{}, my_errors.ErrorNotExists)
	serv := servs.NewUserService(repo)

	ans, err := serv.Add(params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestUserServiceUpdate(t *testing.T) {
	expectedValue := ents.User{
		Id:         1,
		Login:      "user",
		Password:   "123",
		Balance:    100,
		CharitySum: 1000,
	}
	params := ents.UserAdd{
		Login:    "user",
		Password: "123",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().Update(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("user").Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewUserService(repo)

	ans, err := serv.Update("1", params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestUserServiceGetAll(t *testing.T) {
	expectedValue := []ents.User{
		{
			Id:         1,
			Login:      "user",
			Password:   "123",
			Balance:    100,
			CharitySum: 1000,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().Select().Return(expectedValue, nil)
	serv := servs.NewUserService(repo)
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal . expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestUserServiceGetById(t *testing.T) {
	expectedValue := ents.User{
		Id:         1,
		Login:      "user",
		Password:   "123",
		Balance:    100,
		CharitySum: 1000,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewUserService(repo)
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if ans != expectedValue {
		t.Errorf("Should be equal . expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestUserServiceGetByLogin(t *testing.T) {
	expectedValue := ents.User{
		Id:         0,
		Login:      "user",
		Password:   "123",
		Balance:    100,
		CharitySum: 1000,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().SelectByLogin("user").Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("user").Return(expectedValue, nil)
	serv := servs.NewUserService(repo)
	ans, err := serv.GetByLogin("user")
	if err != nil {
		t.Errorf("Some error : %s\n", err)
	}
	if ans != expectedValue {
		t.Errorf("Should be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestUserServiceDonate(t *testing.T) {
	expectedValue := ents.User{
		Id:         1,
		Login:      "user",
		Password:   "123",
		Balance:    110,
		CharitySum: 1000,
	}
	expectedDonatedValue := ents.User{
		Id:         1,
		Login:      "user",
		Password:   "123",
		Balance:    10,
		CharitySum: 1100,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIUserRepository(ctrl)
	repo.EXPECT().Update(expectedDonatedValue).Return(expectedDonatedValue, nil)
	serv := servs.NewUserService(repo)
	err := serv.Donate(&expectedValue, 100)
	if err != nil {
		t.Errorf("Some error at Donate: %s\n", err)
	}

}
