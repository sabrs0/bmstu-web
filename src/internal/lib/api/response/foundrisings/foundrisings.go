package foundrisings

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type GetAllResponse struct {
	Foundrisings []ents.Foundrising `json:"foundrisings"`
}
type BaseResponse struct {
	ents.Foundrising
}
