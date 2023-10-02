package dataaccessTest

import (
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/business/entities"
	repos "github.com/sabrs0/bmstu-web/src/dataAccess/repositories"
)

func TestFoundationInsert(t *testing.T) {
	f := ents.Foundation{
		Id: 1000,
	}
	repo := repos.NewFoundationRepository(mockDB)
	err := repo.Insert(f)
	if err != nil {
		t.Errorf("Cannot insert, some trouble: %s\n", err)
	}
}
func TestFoundationUpdate(t *testing.T) {
	f := ents.Foundation{
		Id:      2000,
		Country: "USA",
	}
	repo := repos.NewFoundationRepository(mockDB)
	err := repo.Update(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.Foundation{
		Id:      1000,
		Country: "USA",
	}
	err = repo.Update(f)
	if err != nil {
		t.Errorf("Cannot update, some trouble: %s\n", err)
	}
}
func TestFoundationDelete(t *testing.T) {
	f := ents.Foundation{
		Id:      2000,
		Country: "USA",
	}
	repo := repos.NewFoundationRepository(mockDB)
	err := repo.Delete(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.Foundation{
		Id:      1000,
		Country: "USA",
	}
	err = repo.Delete(f)
	if err != nil {
		t.Errorf("Cannot delete, some trouble: %s\n", err)
	}
}
func TestFoundationSelect(t *testing.T) {

	repo := repos.NewFoundationRepository(mockDB)
	_, isOk := repo.Select()
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestFoundationSelectById(t *testing.T) {
	repo := repos.NewFoundationRepository(mockDB)
	_, isOk := repo.SelectById(1000)
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestFoundationSelectByLogin(t *testing.T) {

	repo := repos.NewFoundationRepository(mockDB)
	_, isOk := repo.SelectByLogin("USA123")
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestFoundationSelectByName(t *testing.T) {
	repo := repos.NewFoundationRepository(mockDB)
	_, isOk := repo.SelectByName("Foundation of the United States")
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestFoundationSelectByCountry(t *testing.T) {
	repo := repos.NewFoundationRepository(mockDB)
	ans, isOk := repo.SelectByCountry("USA")
	if (!isOk && len(ans) != 0) || (isOk && len(ans) == 0) {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
