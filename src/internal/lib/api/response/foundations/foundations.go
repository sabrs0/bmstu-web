package foundations

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response FoundationsGetAllResponse
type GetAllResponse struct {
	//in: body
	Foundations []ents.Foundation `json:"foundations"`
}

// swagger:response FoundationsBaseResponse
type BaseResponse struct {
	//in: body
	Foundation ents.Foundation
}

// swagger:response FoundationsDonateResponse
type DonateResponse struct {
	//in: body
	Transaction ents.Transaction
}

// swagger:response FoundationsFoundrisingsResponse
type FoundrisingsResponse struct {
	//in: body
	Foundrisings []ents.Foundrising `json:"foundrisings"`
}
