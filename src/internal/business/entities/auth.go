package entities

type Params struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// swagger:parameters  LoginPost
type ParamsRequest struct {
	//in:body
	Params Params
}
