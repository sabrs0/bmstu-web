package foundations

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response FoundationsGetAllResponse
type GetAllResponse struct {
	//in: body
	Foundations []ents.FoundationTransfer `json:"foundations"`
}

// swagger:response FoundationsBaseResponse
type BaseResponse struct {
	//in: body
	Foundation ents.FoundationTransfer
}
type ExtResponse struct {
	Foundation ents.FoundationTransferExtended
}

// swagger:response FoundationsDonateResponse
type DonateResponse struct {
	//in: body
	Transaction ents.TransactionTransfer
}

// swagger:response FoundationsFoundrisingsResponse
type FoundrisingsResponse struct {
	//in: body
	Foundrisings []ents.FoundrisingTransfer `json:"foundrisings"`
}
