package users

import (
	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// swagger:response UsersGetAllResponse
type GetAllResponse struct {
	//in: body
	Users []ents.UserTransfer `json:"users"`
}

// swagger:response UsersBaseResponse
type BaseResponse struct {
	//in: body
	User ents.UserTransfer
}

// swagger:response UsersDonateResponse
type DonateResponse struct {
	//in: body
	Transaction ents.TransactionTransfer
}
