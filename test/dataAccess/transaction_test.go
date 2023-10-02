package dataaccessTest

import (
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	repos "github.com/sabrs0/bmstu-web/src/dataAccess/repositories"
)

func TestTransactionInsert(t *testing.T) {
	f := ents.Transaction{
		Id: 1000,
	}
	repo := repos.NewTransactionRepository(mockDB)
	err := repo.Insert(f)
	if err != nil {
		t.Errorf("Cannot insert, some trouble: %s\n", err)
	}
}
func TestTransactionDelete(t *testing.T) {
	f := ents.Transaction{
		Id: 2000,
	}
	repo := repos.NewTransactionRepository(mockDB)
	err := repo.Delete(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.Transaction{
		Id: 1000,
	}
	err = repo.Delete(f)
	if err != nil {
		t.Errorf("Cannot delete, some trouble: %s\n", err)
	}
}
func TestTransactionSelect(t *testing.T) {

	repo := repos.NewTransactionRepository(mockDB)
	_, isOk := repo.Select()
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestTransactionSelectById(t *testing.T) {
	repo := repos.NewTransactionRepository(mockDB)
	_, isOk := repo.SelectById(1000)
	if !isOk {
		t.Errorf("Cannot Select by id, some trouble\n")
	}
}
func TestTransactionSelectFrom(t *testing.T) {

	repo := repos.NewTransactionRepository(mockDB)
	ans, isOk := repo.SelectFrom(true, 1000)
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select from, some trouble\n")
	}
}
func TestTransactionSelectTo(t *testing.T) {
	repo := repos.NewTransactionRepository(mockDB)
	ans, isOk := repo.SelectTo(true, 1000)
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select to, some trouble\n")
	}
}
func TestTransactionSelectFoundrisingPhilantropIds(t *testing.T) {
	repo := repos.NewTransactionRepository(mockDB)
	ans, isOk := repo.SelectFoundrisingPhilantropIds(1000)
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select foundrising philantrops, some trouble\n")
	}
}
