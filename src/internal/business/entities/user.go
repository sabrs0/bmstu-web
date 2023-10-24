package entities

type UserDonate struct {
	EntityType bool   `json:"entity_type"`
	EntityID   string `json:"entity_id"`
	SumOfMoney string `json:"sum_of_money"`
	Comm       string `json:"comment"`
}
type UserAdd struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type User struct {
	Id               uint64  `gorm:"primaryKey;not null" json:"id"`
	Login            string  `gorm:"not null" json:"login"`
	Password         string  `gorm:"not null" json:"password"`
	Balance          float64 `json:"balance"`
	CharitySum       float64 `json:"charity_sum"`
	ClosedFingAmount uint64  `json:"closed_fing_amount"`
}

func NewUser() User {
	return User{}
}

func NewUserPtr() *User {
	return &User{}
}

func (U *User) SetLogin(newName string) {
	U.Login = newName
}

func (U *User) SetPassword(newPass string) {
	U.Password = newPass
}
