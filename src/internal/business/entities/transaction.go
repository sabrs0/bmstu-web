package entities

type TransactionAdd struct {
	From_essence_type bool    `json:"from_essence_type"`
	From_id           uint64  `gorm:"not null" json:"from_id"`
	To_essence_type   bool    `json:"to_essence_type"`
	Sum_of_money      float64 `json:"sum_of_money"`
	Comment           string  `json:"comment"`
	To_id             uint64  `gorm:"not null" json:"to_id"`
}
type Transaction struct {
	Id uint64 `gorm:"primaryKey;not null" json:"id"`
	TransactionAdd
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
