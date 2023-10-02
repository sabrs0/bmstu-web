package businessLogicTest

import (
	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	"gorm.io/gorm"
)

type FoundationRepMock struct {
	mockDB []ents.Foundation
}

func NewFoundationRepMock(mockDB []ents.Foundation) *FoundationRepMock {
	return &FoundationRepMock{
		mockDB: mockDB,
	}
}
func (mock *FoundationRepMock) Insert(F ents.Foundation) error {
	mock.mockDB = append(mock.mockDB, F)
	return nil
}
func (mock *FoundationRepMock) Update(F ents.Foundation) error {
	for i, f := range mock.mockDB {
		if f.Id == F.Id {
			mock.mockDB[i] = F
			return nil
		}
	}
	return nil
}
func (mock *FoundationRepMock) Delete(F ents.Foundation) error {
	for i, f := range mock.mockDB {
		if f.Id == F.Id {
			mock.mockDB[i], mock.mockDB[len(mock.mockDB)-1] = mock.mockDB[len(mock.mockDB)-1], mock.mockDB[i]
			mock.mockDB = mock.mockDB[:len(mock.mockDB)-1]
			return nil
		}
	}
	return nil
}
func (mock *FoundationRepMock) Select() ([]ents.Foundation, bool) {
	return mock.mockDB, true
}
func (mock *FoundationRepMock) SelectById(id uint64) (ents.Foundation, bool) {
	for _, f := range mock.mockDB {
		if f.Id == id {
			return f, true
		}
	}
	return ents.NewFoundation(), false
}
func (mock *FoundationRepMock) SelectByLogin(name string) (ents.Foundation, bool) {
	for _, f := range mock.mockDB {
		if f.Login == name {
			return f, true
		}
	}
	return ents.NewFoundation(), false
}
func (mock *FoundationRepMock) SelectByName(name string) (ents.Foundation, bool) {
	for _, f := range mock.mockDB {
		if f.Name == name {
			return f, true
		}
	}
	return ents.NewFoundation(), false
}
func (mock *FoundationRepMock) SelectByCountry(country string) ([]ents.Foundation, bool) {
	ans := []ents.Foundation{}
	for _, f := range mock.mockDB {
		if f.Country == country {
			ans = append(ans, f)
		}
	}
	return ans, false
}
func (mock *FoundationRepMock) GetDB() *gorm.DB {
	return nil
}

type FoundrisingRepMock struct {
	mockDB []ents.Foundrising
}

func NewFoundrisingRepMock(mockDB []ents.Foundrising) *FoundrisingRepMock {
	return &FoundrisingRepMock{
		mockDB: mockDB,
	}
}
func (mock *FoundrisingRepMock) Insert(F ents.Foundrising) error {
	mock.mockDB = append(mock.mockDB, F)
	return nil
}
func (mock *FoundrisingRepMock) Update(F ents.Foundrising) error {
	for i, f := range mock.mockDB {
		if f.Id == F.Id {
			mock.mockDB[i] = F
			return nil
		}
	}
	return nil
}
func (mock *FoundrisingRepMock) Delete(F ents.Foundrising) error {
	for i, f := range mock.mockDB {
		if f.Id == F.Id {
			mock.mockDB[i], mock.mockDB[len(mock.mockDB)-1] = mock.mockDB[len(mock.mockDB)-1], mock.mockDB[i]
			mock.mockDB = mock.mockDB[:len(mock.mockDB)-1]
			return nil
		}
	}
	return nil
}
func (mock *FoundrisingRepMock) Select() ([]ents.Foundrising, bool) {
	return mock.mockDB, true
}
func (mock *FoundrisingRepMock) SelectById(id uint64) (ents.Foundrising, bool) {
	for _, f := range mock.mockDB {
		if f.Id == id {
			return f, true
		}
	}
	return ents.NewFoundrising(), false
}
func (mock *FoundrisingRepMock) SelectByFoundId(id uint64) ([]ents.Foundrising, bool) {
	ans := []ents.Foundrising{}
	for _, f := range mock.mockDB {
		if f.Found_id == id {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *FoundrisingRepMock) SelectByCreateDate(date string) ([]ents.Foundrising, bool) {
	ans := []ents.Foundrising{}
	for _, f := range mock.mockDB {
		if f.Creation_date == date {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *FoundrisingRepMock) SelectByFoundIdActive(id uint64) ([]ents.Foundrising, bool) {
	ans := []ents.Foundrising{}
	for _, f := range mock.mockDB {
		if f.Closing_date.Valid == false && f.Id == id {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *FoundrisingRepMock) SelectByCloseDate(date string) ([]ents.Foundrising, bool) {
	ans := []ents.Foundrising{}
	for _, f := range mock.mockDB {
		if f.Closing_date.String == date {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *FoundrisingRepMock) SelectByIdAndFoundId(id_ uint64, found_id_ uint64) (ents.Foundrising, bool) {
	ans := []ents.Foundrising{}
	for _, f := range mock.mockDB {
		if f.Id == id_ && f.Found_id == found_id_ {
			ans = append(ans, f)
		}
	}
	return ans[0], true
}
func (mock *FoundrisingRepMock) GetDB() *gorm.DB {
	return nil
}

type TransactionRepMock struct {
	mockDB []ents.Transaction
}

func NewTransactionRepMock(mockDB []ents.Transaction) *TransactionRepMock {
	return &TransactionRepMock{
		mockDB: mockDB,
	}
}
func (mock *TransactionRepMock) Insert(T ents.Transaction) error {
	mock.mockDB = append(mock.mockDB, T)
	return nil
}
func (mock *TransactionRepMock) Delete(T ents.Transaction) error {
	for i, f := range mock.mockDB {
		if f.Id == T.Id {
			mock.mockDB[i], mock.mockDB[len(mock.mockDB)-1] = mock.mockDB[len(mock.mockDB)-1], mock.mockDB[i]
			mock.mockDB = mock.mockDB[:len(mock.mockDB)-1]
			return nil
		}
	}
	return nil
}
func (mock *TransactionRepMock) Select() ([]ents.Transaction, bool) {
	return mock.mockDB, true
}
func (mock *TransactionRepMock) SelectFrom(type_ bool, id uint64) ([]ents.Transaction, bool) {
	ans := []ents.Transaction{}
	for _, f := range mock.mockDB {
		if f.From_essence_type == type_ && f.From_id == id {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *TransactionRepMock) SelectTo(type_ bool, id uint64) ([]ents.Transaction, bool) {
	ans := []ents.Transaction{}
	for _, f := range mock.mockDB {
		if f.To_essence_type == type_ && f.To_id == id {
			ans = append(ans, f)
		}
	}
	return ans, true
}
func (mock *TransactionRepMock) SelectById(id uint64) (ents.Transaction, bool) {
	for _, f := range mock.mockDB {
		if f.Id == id {
			return f, true
		}
	}
	return ents.NewTransaction(), false
}

func (mock *TransactionRepMock) SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, bool) {
	ans := []uint64{}
	for _, f := range mock.mockDB {
		if f.From_essence_type == true && f.To_essence_type == true && f.To_id == foundrising_id {
			ans = append(ans, f.From_id)
		}
	}
	return ans, true
}
func (mock *TransactionRepMock) GetDB() *gorm.DB {
	return nil
}

type UserRepMock struct {
	mockDB []ents.User
}

func NewUserRepMock(mockDB []ents.User) *UserRepMock {
	return &UserRepMock{
		mockDB: mockDB,
	}
}
func (mock *UserRepMock) Insert(U ents.User) error {
	mock.mockDB = append(mock.mockDB, U)
	return nil
}
func (mock *UserRepMock) Update(U ents.User) error {
	for i, f := range mock.mockDB {
		if f.Id == U.Id {
			mock.mockDB[i] = U
			return nil
		}
	}
	return nil
}
func (mock *UserRepMock) Delete(U ents.User) error {
	for i, f := range mock.mockDB {
		if f.Id == U.Id {
			mock.mockDB[i], mock.mockDB[len(mock.mockDB)-1] = mock.mockDB[len(mock.mockDB)-1], mock.mockDB[i]
			mock.mockDB = mock.mockDB[:len(mock.mockDB)-1]
			return nil
		}
	}
	return nil
}
func (mock *UserRepMock) Select() ([]ents.User, bool) {
	return mock.mockDB, true
}
func (mock *UserRepMock) SelectById(id uint64) (ents.User, bool) {
	for _, f := range mock.mockDB {
		if f.Id == id {
			return f, true
		}
	}
	return ents.NewUser(), false
}
func (mock *UserRepMock) SelectByLogin(name string) (ents.User, bool) {
	for _, f := range mock.mockDB {
		if f.Login == name {
			return f, true
		}
	}
	return ents.NewUser(), false
}
func (mock *UserRepMock) GetDB() *gorm.DB {
	return nil
}
