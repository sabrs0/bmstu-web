package foundations

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type GetAllResponse struct {
	Foundations []ents.Foundation `json:"foundations"`
}
type BaseResponse struct {
	ents.Foundation
}
type DonateResponse struct {
	ents.Transaction
}
type FoundrisingsResponse struct {
	Foundrisings []ents.Foundrising `json:"foundrisings"`
}
