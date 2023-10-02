package businessLogicTest

import (
	"reflect"
	"testing"

	chk "github.com/sabrs0/bmstu-web/src/business/checker"
	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/business/services"
)

func TestTransactionServiceAdd(t *testing.T) {
	var Transactions []ents.Transaction = []ents.Transaction{
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
	serv := servs.NewTransactionService(NewTransactionRepMock(Transactions))
	pars := chk.NewTransactionMainParams(true, 2, false, 100.00, "transact", 1)
	err := serv.Add(pars)
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestTransactionServiceDelete(t *testing.T) {
	var Transactions []ents.Transaction = []ents.Transaction{
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
	serv := servs.NewTransactionService(NewTransactionRepMock(Transactions))
	err := serv.Delete("1")
	if err != nil {
		t.Errorf("Shoud be nil, but err is %s\n", err)
	}
}
func TestTransactionServiceGetAll(t *testing.T) {
	var Transactions []ents.Transaction = []ents.Transaction{
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
	serv := servs.NewTransactionService(NewTransactionRepMock(Transactions))
	ans, err := serv.GetAll()
	if err != nil {
		t.Errorf("Some error at getAllTransactions: %s\n", err)
	}
	if !reflect.DeepEqual(Transactions, ans) {
		t.Errorf("Shoud be equal getAll. ans = %v, src = %v\n", ans, Transactions)
	}
}
func TestTransactionServiceGetById(t *testing.T) {
	var Transactions []ents.Transaction = []ents.Transaction{
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
	serv := servs.NewTransactionService(NewTransactionRepMock(Transactions))
	ans, err := serv.GetById("1")
	if err != nil {
		t.Errorf("Some error at getById: %s\n", err)
	}
	if ans != Transactions[0] {
		t.Errorf("Should be equal getById. ans = %v, src = %v\n", ans, Transactions)
	}
}
