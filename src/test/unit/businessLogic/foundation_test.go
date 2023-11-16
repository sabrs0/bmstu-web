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

func TestFoundationServiceAdd(t *testing.T) {

	params := ents.FoundationAdd{

		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	expectedValue := ents.Foundation{
		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().Insert(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("Foundation124").Return(ents.Foundation{}, my_errors.ErrorNotExists)
	serv := servs.NewFoundationService(repo)

	ans, err := serv.Add(params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestFoundationServiceUpdate(t *testing.T) {
	params := ents.FoundationAdd{
		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	expectedValue := ents.Foundation{
		Id:       1,
		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().Update(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("Foundation124").Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewFoundationService(repo)

	ans, err := serv.Update("1", params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestFoundationServiceGetAll(t *testing.T) {

	expectedValue := []ents.Foundation{
		{
			Name:     "F",
			Password: "123",
			Country:  "Россия",
			Login:    "Foundation124",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().Select().Return(expectedValue, nil)
	serv := servs.NewFoundationService(repo)
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal getAll. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestFoundationServiceGetById(t *testing.T) {
	expectedValue := ents.Foundation{
		Id:       1,
		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewFoundationService(repo)
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if ans != expectedValue {
		t.Errorf("Should be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestFoundationServiceGetBylogin(t *testing.T) {
	expectedValue := ents.Foundation{
		Id:       1,
		Name:     "F",
		Password: "123",
		Country:  "Россия",
		Login:    "Foundation124",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().SelectByLogin("Foundation124").Return(expectedValue, nil)
	repo.EXPECT().SelectByLogin("Foundation124").Return(expectedValue, nil)

	serv := servs.NewFoundationService(repo)
	ans, err := serv.GetByLogin("Foundation124")
	if err != nil {
		t.Errorf("Some error : %s\n", err)
	}
	if ans != expectedValue {
		t.Errorf("Should be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestFoundationServiceDonate(t *testing.T) {
	expectedValue := ents.Foundation{
		Id:           1,
		Name:         "Foundation",
		Password:     "123",
		Country:      "Россия",
		Login:        "Foundation123",
		Fund_balance: 110.00,
	}
	expectedDonatedValue := ents.Foundation{
		Id:              1,
		Name:            "Foundation",
		Password:        "123",
		Country:         "Россия",
		Login:           "Foundation123",
		Fund_balance:    60.00,
		Outcome_history: 50.00,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundationRepository(ctrl)
	repo.EXPECT().Update(expectedDonatedValue).Return(expectedDonatedValue, nil)
	serv := servs.NewFoundationService(repo)
	err := serv.Donate(&expectedValue, 50.00)
	if err != nil {
		t.Errorf("Some error at Donate: %s\n", err)
	}
}
