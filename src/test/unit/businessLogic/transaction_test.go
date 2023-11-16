package businessLogicTest

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/test/unit/businessLogic/mocks"
)

func TestTransactionServiceAdd(t *testing.T) {
	expectedValue := ents.Transaction{
		Id:                0,
		From_essence_type: true,
		From_id:           1,
		To_essence_type:   true,
		Sum_of_money:      100.00,
		Comment:           "Transacting",
		To_id:             1,
	}
	params := ents.TransactionAdd{
		From_essence_type: true,
		From_id:           1,
		To_essence_type:   true,
		Sum_of_money:      100.00,
		Comment:           "Transacting",
		To_id:             1,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockITransactionRepository(ctrl)
	repo.EXPECT().Insert(expectedValue).Return(expectedValue, nil)
	serv := servs.NewTransactionService(repo)
	ans, err := serv.Add(params)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}

}
func TestTransactionServiceDelete(t *testing.T) {
	expectedValue := ents.Transaction{
		Id:                1,
		From_essence_type: true,
		From_id:           1,
		To_essence_type:   true,
		Sum_of_money:      100.00,
		Comment:           "Transacting",
		To_id:             1,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockITransactionRepository(ctrl)
	repo.EXPECT().Delete(expectedValue).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	serv := servs.NewTransactionService(repo)
	ans, err := serv.Delete("1")
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal values, but:\n expected = %#v\nrecieved = %#v\n", expectedValue, ans)
	}
}
func TestTransactionServiceGetAll(t *testing.T) {
	expectedValue := []ents.Transaction{
		{
			Id:                1,
			From_essence_type: true,
			From_id:           1,
			To_essence_type:   true,
			Sum_of_money:      100.00,
			Comment:           "Transacting",
			To_id:             1,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockITransactionRepository(ctrl)
	repo.EXPECT().Select().Return(expectedValue, nil)
	serv := servs.NewTransactionService(repo)
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error at getAllTransactions: %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal getAll. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
func TestTransactionServiceGetById(t *testing.T) {
	expectedValue := ents.Transaction{
		Id:                1,
		From_essence_type: true,
		From_id:           1,
		To_essence_type:   true,
		Sum_of_money:      100.00,
		Comment:           "Transacting",
		To_id:             1,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockITransactionRepository(ctrl)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)
	repo.EXPECT().SelectById(uint64(1)).Return(expectedValue, nil)

	serv := servs.NewTransactionService(repo)
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error at getAllFoudnations: %s\n", err)
	}
	if !reflect.DeepEqual(expectedValue, ans) {
		t.Errorf("Shoud be equal getById. expected = %#v\nrecieved = %#v\n\n", expectedValue, ans)
	}
}
