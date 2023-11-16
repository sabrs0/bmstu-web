package postgres

import (
	"database/sql"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db_ *sql.DB) *UserRepository {
	return &UserRepository{DB: db_}
}
func (FR *UserRepository) Insert(U ents.User) (ents.User, error) {
	var LastId uint64
	row := FR.DB.QueryRow("SELECT id from User_tab ORDER BY id DESC LIMIT 1")
	err := row.Scan(&LastId)
	if err != nil {
		if err != sql.ErrNoRows {
			return ents.User{}, err
		}
	}
	U.Id = LastId + 1
	_, err = FR.DB.Exec(`insert into User_tab(
		id, 
		login, 
		password, 
		balance, 
		charity_Sum,
		closed_Fing_Amount) 
		values ($1, $2, $3, $4, $5, $6)`,
		U.Id,
		U.Login,
		U.Password,
		U.Balance,
		U.CharitySum,
		U.ClosedFingAmount,
	)
	if err != nil {
		return ents.User{}, fmt.Errorf("error in insert User repo: %s", err.Error())
	}
	return U, nil
}

func (FR *UserRepository) Delete(F ents.User) (ents.User, error) {
	_, err := FR.DB.Exec("delete from User_tab where id = $1", F.Id)
	if err != nil {
		return ents.User{}, fmt.Errorf("error in Delete User repo: %s", err.Error())
	}
	return F, nil
}

func (FR *UserRepository) Update(U ents.User) (ents.User, error) {
	_, err := FR.DB.Exec(`update User_tab set 
	login = $1, 
	password = $2, 
	balance = $3, 
	charity_Sum = $4,
	closed_Fing_Amount = $5
	where id = $6`, U.Login,
		U.Password,
		U.Balance,
		U.CharitySum,
		U.ClosedFingAmount,
		U.Id)
	if err != nil {
		return ents.User{}, fmt.Errorf("error in Update User_tab repo: %s", err.Error())
	}
	return U, nil
}

func (FR *UserRepository) Select() ([]ents.User, error) {
	var User_tab []ents.User
	//result := FR.DB.Table("User_tab").Order("ID").Find(&User_tab)
	rows, err := FR.DB.Query("select * from User_tab")
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

	}
	defer rows.Close()
	for rows.Next() {
		U := ents.User{}
		err := rows.Scan(&U.Id,
			&U.Login,
			&U.Password,
			&U.Balance,
			&U.CharitySum,
			&U.ClosedFingAmount,
		)
		if err != nil {
			return nil, err
		}
		User_tab = append(User_tab, U)
	}

	return User_tab, nil
}
func (FR *UserRepository) SelectById(id uint64) (ents.User, error) {
	var U ents.User
	row := FR.DB.QueryRow("select * from User_tab where id = $1", id)
	err := row.Scan(&U.Id,
		&U.Login,
		&U.Password,
		&U.Balance,
		&U.CharitySum,
		&U.ClosedFingAmount)
	if err != nil {
		return ents.User{}, err
	}
	return U, nil
}

func (FR *UserRepository) SelectByLogin(name string) (ents.User, error) {
	var U ents.User
	row := FR.DB.QueryRow("select * from User_tab where login = $1", name)
	err := row.Scan(&U.Id,
		&U.Login,
		&U.Password,
		&U.Balance,
		&U.CharitySum,
		&U.ClosedFingAmount)
	if err != nil {
		return ents.User{}, err
	}
	return U, nil
}
