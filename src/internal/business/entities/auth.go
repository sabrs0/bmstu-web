package entities

// swagger:model LoginParams
type Params struct {
	// in: query
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// swagger:model Token
type Token struct {
	Token string `json:"token"`
}
