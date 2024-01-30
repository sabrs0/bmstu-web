package entities

import "strconv"

type TransactionAdd struct {
	From_essence_type bool    `json:"from_essence_type"`
	From_id           uint64  `gorm:"not null" json:"from_id"`
	To_essence_type   bool    `json:"to_essence_type"`
	Sum_of_money      float64 `json:"sum_of_money"`
	Comment           string  `json:"comment"`
	To_id             uint64  `gorm:"not null" json:"to_id"`
}

type Transaction struct {
	Id                uint64  `gorm:"primaryKey;not null" json:"id"`
	From_essence_type bool    `json:"from_essence_type"`
	From_id           uint64  `gorm:"not null" json:"from_id"`
	To_essence_type   bool    `json:"to_essence_type"`
	Sum_of_money      float64 `json:"sum_of_money"`
	Comment           string  `json:"comment"`
	To_id             uint64  `gorm:"not null" json:"to_id"`
}

// swagger:model
type TransactionTransfer struct {
	Id           uint64  `json:"id"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Sum_of_money float64 `json:"sum_of_money"`
	Comment      string  `json:"comment"`
}

func NewTransactionTransfer(t Transaction) TransactionTransfer {
	var tp TransactionTransfer = TransactionTransfer{
		Id:           t.Id,
		Sum_of_money: t.Sum_of_money,
		Comment:      t.Comment,
	}
	if t.From_essence_type == FROM_FOUNDATION {
		tp.From = "Foundation"
	} else {
		tp.From = "User"
	}
	tp.From = tp.From + " " + strconv.Itoa(int(t.From_id))
	if t.To_essence_type == TO_FOUNDATION {
		tp.To = "Foundation"
	} else {
		tp.To = "Foundrising"
	}
	tp.To = tp.To + " " + strconv.Itoa(int(t.To_id))
	return tp
}

func NewTransaction() Transaction {
	return Transaction{}
}

func NewTransactionPtr() *Transaction {
	return &Transaction{}
}

func (T *Transaction) SetFromType(type_ bool) {
	T.From_essence_type = type_
}

func (T *Transaction) SetToType(type_ bool) {
	T.To_essence_type = type_
}

func (T *Transaction) SetFromId(id_ uint64) {
	T.From_id = id_
}

func (T *Transaction) SetToId(id_ uint64) {
	T.To_id = id_
}

func (T *Transaction) SetSumOfMoney(sum float64) {
	T.Sum_of_money = sum
}

func (T *Transaction) SetComment(newName string) {
	T.Comment = newName
}

const (
	FROM_USER       = false
	FROM_FOUNDATION = true
	TO_FOUNDATION   = false
	TO_FOUNDRISING  = true
)
