package auth

type TokenResponse struct {
	Token string `json:"token"`
}

// swagger:response LoginResponse
type LoginResponse struct {
	//in:body
	TokenResponse TokenResponse
}
