package entities

import "database/sql"

const DateFormat string = "2006-01-02"

type FoundrisingPut struct {
	Description  string `json:"description"`
	Required_sum string `json:"required_sum"`
}
type FoundrisingAdd struct {
	Found_id      uint64 `json:"found_id"`
	Description   string `json:"description"`
	Required_sum  string `json:"required_sum"`
	Creation_date string `json:"creation_date"`
}
type Foundrising struct {
	Id                 uint64         `gorm:"primaryKey;not null" json:"id"`
	Found_id           uint64         `gorm:"not null" json:"found_id"`
	Description        string         `json:"description"`
	Required_sum       float64        `json:"required_sum"`
	Current_sum        float64        `json:"current_sum"`
	Philantrops_amount uint64         `json:"philantrops_amount"`
	Creation_date      string         `json:"creation_date"`
	Closing_date       sql.NullString `json:"closing_date"`
}

func NewFoundrising() Foundrising {
	return Foundrising{}
}

func NewFoundrisingPtr() *Foundrising {
	return &Foundrising{}
}

func (F *Foundrising) SetDescription(n string) {
	F.Description = n
}

func (F *Foundrising) SetReqSum(sum float64) {
	F.Required_sum = sum
}
func (F *Foundrising) SetCreateDate(newName string) {
	F.Creation_date = newName
}

func (F *Foundrising) SetFoundId(id uint64) {
	F.Found_id = id
}
