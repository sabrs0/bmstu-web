package transactions

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type GetAllResponse struct {
	Transactions []ents.Transaction `json:"transactions"`
}
type BaseResponse struct {
	ents.Transaction
}
