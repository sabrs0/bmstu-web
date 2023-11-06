package businessLogicTest

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/test/businessLogic/mocks"
)

func TestFoundrisingServiceAdd(t *testing.T) {
	expectedValue := ents.Foundrising{
		Id:            0,
		Found_id:      1,
		Description:   "123",
		Required_sum:  100,
		Creation_date: "01-01-2002",
		Closing_date:  sql.NullString{Valid: false},
	}
	params := ents.FoundrisingAdd{
		Found_id:      1,
		Description:   "123",
		Required_sum:  "100",
		Creation_date: "01-01-2002",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().Insert(expectedValue).Return(expectedValue, nil)
	serv := servs.NewFoundrisingService(repo)

	ans, err := serv.Add(params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}

}
func TestFoundrisingServiceUpdate(t *testing.T) {
	expectedValue := ents.Foundrising{
		Id:            1,
		Found_id:      1,
		Description:   "123",
		Required_sum:  100,
		Creation_date: "01-01-2021",
		Closing_date:  sql.NullString{String: "01-01-2021"},
	}
	params := ents.FoundrisingPut{
		Description:  "123",
		Required_sum: "100",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().Update(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewFoundrisingService(repo)

	ans, err := serv.Update("1", params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestFoundrisingServiceGetAll(t *testing.T) {
	expectedValue := []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().Select().Return(expectedValue, nil)
	serv := servs.NewFoundrisingService(repo)
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal getAll. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}

}
func TestFoundrisingServiceGetById(t *testing.T) {
	expectedValue := ents.Foundrising{
		Id:            1,
		Found_id:      1,
		Description:   "123",
		Required_sum:  100,
		Creation_date: "01-01-2021",
		Closing_date:  sql.NullString{String: "01-01-2021"},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewFoundrisingService(repo)
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if ans != expectedValue {
		t.Errorf("Should be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}

}
func TestFoundrisingServiceGetByCreateDate(t *testing.T) {
	expectedValue := []ents.Foundrising{
		{
			Id:            1,
			Found_id:      1,
			Description:   "123",
			Required_sum:  100,
			Creation_date: "01-01-2021",
			Closing_date:  sql.NullString{String: "01-01-2021"},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().SelectByCreateDate("01-01-2021").Return(expectedValue, nil)
	serv := servs.NewFoundrisingService(repo)
	ans, err := serv.GetByCreateDate("01-01-2021")
	if err != nil {
		t.Errorf("Some error: %s\n", err)
	}
	if !reflect.DeepEqual(ans, expectedValue) {
		t.Errorf("Should be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestFoundrisingServiceAcceptDonate(t *testing.T) {
	expectedValue := ents.Foundrising{
		Id:            1,
		Found_id:      1,
		Description:   "123",
		Required_sum:  100,
		Creation_date: "01-01-2021",
		Closing_date:  sql.NullString{Valid: false},
	}
	expectedAcceptedValue := ents.Foundrising{
		Id:            1,
		Found_id:      1,
		Description:   "123",
		Current_sum:   50,
		Required_sum:  100,
		Creation_date: "01-01-2021",
		Closing_date:  sql.NullString{Valid: false},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockIFoundrisingRepository(ctrl)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().Update(expectedAcceptedValue).Return(expectedAcceptedValue, nil)
	serv := servs.NewFoundrisingService(repo)
	_, err := serv.AcceptDonate("1", 50.00, false)
	if err != nil {
		t.Errorf("Some error : %s\n", err)
	}
}
