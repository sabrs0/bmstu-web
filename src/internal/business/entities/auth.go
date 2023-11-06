package entities

// swagger:model
type Params struct {
	// in: query
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// swagger:model
type Token struct {
	Token string `json:"token"`
}
