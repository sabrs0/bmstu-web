package entities

var Countries [9]string = [9]string{"США", "Россия", "Великобритания", "Канада", "Франция", "Германия", "Китай", "Италия", "Испания"}

type FoundationAdd struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Country  string `json:"country"`
	Login    string `json:"login"`
}

type FoundationDonate struct {
	Foundrising_id string `json:"foundrising_id"`
	Sum_of_Money   string `json:"sum_of_money"`
	Comm           string `json:"comment"`
}
type FoundationReplenish struct {
	Sum_of_Money string `json:"sum_of_money"`
}

// swagger:model
type Foundation struct {
	Id                      uint64  `gorm:"primaryKey;not null" json:"id"`
	Name                    string  `gorm:"not null" json:"name"`
	Password                string  `gorm:"not null" json:"password"`
	CurFoudrisingAmount     uint32  `json:"cur_foudrising_amount"`
	ClosedFoundrisingAmount uint32  `json:"closed_foundrising_amount"`
	Fund_balance            float64 `json:"fund_balance"`
	Income_history          float64 `json:"income_history"`
	Outcome_history         float64 `json:"outcome_history"`
	Volunteer_amount        uint32  `json:"volunteer_amount"`
	Country                 string  `json:"country"`
	Login                   string  `gorm:"not null" json:"login"`
}

// swagger:model
type FoundationTransfer struct {
	Id                      uint64 `json:"id"`
	Name                    string `json:"name"`
	CurFoudrisingAmount     uint32 `json:"cur_foudrising_amount"`
	ClosedFoundrisingAmount uint32 `json:"closed_foundrising_amount"`
	Volunteer_amount        uint32 `json:"volunteer_amount"`
	Country                 string `json:"country"`
}

func NewFoundationTransfer(f Foundation) FoundationTransfer {
	return FoundationTransfer{
		Id:                      f.Id,
		Name:                    f.Name,
		CurFoudrisingAmount:     f.CurFoudrisingAmount,
		ClosedFoundrisingAmount: f.ClosedFoundrisingAmount,
		Volunteer_amount:        f.Volunteer_amount,
		Country:                 f.Country,
	}
}

type FoundationTransferExtended struct {
	Id                      uint64  `json:"id"`
	Name                    string  `json:"name"`
	CurFoudrisingAmount     uint32  `json:"cur_foudrising_amount"`
	ClosedFoundrisingAmount uint32  `json:"closed_foundrising_amount"`
	Fund_balance            float64 `json:"fund_balance"`
	Income_history          float64 `json:"income_history"`
	Outcome_history         float64 `json:"outcome_history"`
	Volunteer_amount        uint32  `json:"volunteer_amount"`
	Country                 string  `json:"country"`
	Login                   string  `json:"login"`
	Password                string  `json:"password"`
}

func NewFoundationTransferExtended(f Foundation) FoundationTransferExtended {
	return FoundationTransferExtended{
		Id:                      f.Id,
		Name:                    f.Name,
		CurFoudrisingAmount:     f.CurFoudrisingAmount,
		ClosedFoundrisingAmount: f.ClosedFoundrisingAmount,
		Volunteer_amount:        f.Volunteer_amount,
		Country:                 f.Country,
		Fund_balance:            f.Fund_balance,
		Income_history:          f.Income_history,
		Outcome_history:         f.Outcome_history,
		Login:                   f.Login,
		Password:                f.Password,
	}
}

func NewFoundation() Foundation {
	return Foundation{}
}

func NewFoundationPtr() *Foundation {
	return &Foundation{}
}

func (F *Foundation) SetLogin(newName string) {
	F.Login = newName
}

func (F *Foundation) SetPassword(newPass string) {
	F.Password = newPass
}

func (F *Foundation) SetName(newName string) {
	F.Name = newName
}

func (F *Foundation) SetCountry(newCntry string) {
	F.Country = newCntry
}

// swagger:parameters FoundationsDonate
type FoundationDonateRequest struct {
	//in: path
	//required: true
	//type: integer
	//format: int64
	Id uint64 `json:"id"`
	//in:body
	Params FoundationDonate
}

// swagger:parameters  FoundationsPost
type FoundationPostRequest struct {
	//in:body
	Params FoundationAdd
}

// swagger:parameters  FoundationsUpdate
type FoundationUpdateRequest struct {
	//required: true
	//in: path
	//type: integer
	//format: int64
	Id uint64 `json:"id"`
	//required: true
	//in:body
	Params FoundationAdd
}
