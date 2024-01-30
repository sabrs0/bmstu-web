package foundrisings

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response FoundrisingsGetAllResponse
type GetAllResponse struct {
	//in: body
	Foundrisings []ents.FoundrisingTransfer `json:"foundrisings"`
}

// swagger:response FoundrisingsBaseResponse
type BaseResponse struct {
	//in: body
	Foundrising ents.FoundrisingTransfer
}
