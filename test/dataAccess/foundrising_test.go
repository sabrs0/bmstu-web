package dataaccessTest

import (
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	repos "github.com/sabrs0/bmstu-web/src/dataAccess/repositories"
)

func TestFoundrisingInsert(t *testing.T) {
	f := ents.Foundrising{
		Id: 1000,
	}
	repo := repos.NewFoundrisingRepository(mockDB)
	err := repo.Insert(f)
	if err != nil {
		t.Errorf("Cannot insert, some trouble: %s\n", err)
	}
}
func TestFoundrisingUpdate(t *testing.T) {
	f := ents.Foundrising{
		Id:          2000,
		Description: "descr",
	}
	repo := repos.NewFoundrisingRepository(mockDB)
	err := repo.Update(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.Foundrising{
		Id:          1000,
		Description: "descr",
	}
	err = repo.Update(f)
	if err != nil {
		t.Errorf("Cannot update, some trouble: %s\n", err)
	}
}
func TestFoundrisingDelete(t *testing.T) {
	f := ents.Foundrising{
		Id:          2000,
		Description: "descr",
	}
	repo := repos.NewFoundrisingRepository(mockDB)
	err := repo.Delete(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.Foundrising{
		Id:          1000,
		Description: "descr",
	}
	err = repo.Delete(f)
	if err != nil {
		t.Errorf("Cannot delete, some trouble: %s\n", err)
	}
}
func TestFoundrisingSelect(t *testing.T) {

	repo := repos.NewFoundrisingRepository(mockDB)
	_, isOk := repo.Select()
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestFoundrisingSelectById(t *testing.T) {
	repo := repos.NewFoundrisingRepository(mockDB)
	_, isOk := repo.SelectById(1000)
	if !isOk {
		t.Errorf("Cannot Select by Id, some trouble\n")
	}
}
func TestFoundrisingSelectByFoundId(t *testing.T) {

	repo := repos.NewFoundrisingRepository(mockDB)
	ans, isOk := repo.SelectByFoundId(1000)
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select by found id, some trouble\n")
	}
}
func TestFoundrisingSelectByCreateDate(t *testing.T) {
	repo := repos.NewFoundrisingRepository(mockDB)
	ans, isOk := repo.SelectByCreateDate("01-01-2021")
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select by create date, some trouble\n")
	}
}
func TestFoundrisingSelectByCloseDate(t *testing.T) {
	repo := repos.NewFoundrisingRepository(mockDB)
	ans, isOk := repo.SelectByCloseDate("01-01-2021")
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select by close date, some trouble\n")
	}
}
func TestFoundrisingSelectByFoundIdActive(t *testing.T) {
	repo := repos.NewFoundrisingRepository(mockDB)
	ans, isOk := repo.SelectByFoundIdActive(1000)
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select by found id active, some trouble\n")
	}
}
func TestFoundrisingSelectByIdAndFoundId(t *testing.T) {
	repo := repos.NewFoundrisingRepository(mockDB)
	_, isOk := repo.SelectByIdAndFoundId(1000, 1000)
	if !isOk {
		t.Errorf("Cannot Select by id and found id, some trouble\n")
	}
}
