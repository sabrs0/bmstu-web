package transactions

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response TransactionsGetAllResponse
type GetAllResponse struct {
	//in: body
	Transactions []ents.TransactionTransfer `json:"transactions"`
}

// swagger:response TransactionsBaseResponse
type BaseResponse struct {
	//in: body
	Transaction ents.TransactionTransfer
}
