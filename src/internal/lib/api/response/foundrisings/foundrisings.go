package foundrisings

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response FoundrisingsGetAllResponse
type GetAllResponse struct {
	//in: body
	Foundrisings []ents.Foundrising `json:"foundrisings"`
}

// swagger:response FoundrisingsBaseResponse
type BaseResponse struct {
	//in: body
	Foundrising ents.Foundrising
}
