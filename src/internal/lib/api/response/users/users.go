package users

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type GetAllResponse struct {
	Users []ents.User `json:"users"`
}
type BaseResponse struct {
	ents.User
}
type DonateResponse struct {
	ents.Transaction
}
