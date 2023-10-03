package dataaccessTest

import (
	"testing"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	repos "github.com/sabrs0/bmstu-web/src/internal/dataAccess/repositories"
)

func TestUserInsert(t *testing.T) {
	f := ents.User{
		Id: 1000,
	}
	repo := repos.NewUserRepository(mockDB)
	err := repo.Insert(f)
	if err != nil {
		t.Errorf("Cannot insert, some trouble: %s\n", err)
	}
}
func TestUserUpdate(t *testing.T) {
	f := ents.User{
		Id:    2000,
		Login: "user",
	}
	repo := repos.NewUserRepository(mockDB)
	err := repo.Update(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.User{
		Id:    1000,
		Login: "user",
	}
	err = repo.Update(f)
	if err != nil {
		t.Errorf("Cannot update, some trouble: %s\n", err)
	}
}
func TestUserDelete(t *testing.T) {
	f := ents.User{
		Id:    2000,
		Login: "user",
	}
	repo := repos.NewUserRepository(mockDB)
	err := repo.Delete(f)
	if err == nil {
		t.Errorf("Need error, but err is nil\n")
	}
	f = ents.User{
		Id:    1000,
		Login: "user",
	}
	err = repo.Delete(f)
	if err != nil {
		t.Errorf("Cannot delete, some trouble: %s\n", err)
	}
}
func TestUserSelect(t *testing.T) {

	repo := repos.NewUserRepository(mockDB)
	_, isOk := repo.Select()
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestUserSelectById(t *testing.T) {
	repo := repos.NewUserRepository(mockDB)
	_, isOk := repo.SelectById(1000)
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
func TestUserSelectByLogin(t *testing.T) {

	repo := repos.NewUserRepository(mockDB)
	_, isOk := repo.SelectByLogin("Login123")
	if !isOk {
		t.Errorf("Cannot Select, some trouble\n")
	}
}
