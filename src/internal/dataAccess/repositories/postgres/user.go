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
func (FR *UserRepository) Insert(U ents.User) error {
	var LastId uint64
	row := FR.DB.QueryRow("SELECT id from User_tab ORDER BY id DESC LIMIT 1")
	err := row.Scan(&LastId)
	if err != nil {
		return err
	}
	_, err = FR.DB.Exec(`insert into User_tab(
		id, 
		login, 
		password, 
		balance, 
		charity_Sum,
		closed_Fing_Amount,) 
		values ($1, $2, $3, $4, $5, $6)`,
		LastId+1,
		U.Login,
		U.Password,
		U.Balance,
		U.CharitySum,
		U.ClosedFingAmount,
	)
	if err != nil {
		return fmt.Errorf("error in insert User repo")
	} else {
		return nil
	}
}

func (FR *UserRepository) Delete(F ents.User) error {
	_, err := FR.DB.Exec("delete from User_tab where id = $1", F.Id)
	if err != nil {
		return fmt.Errorf("error in Delete User repo")
	} else {
		return nil
	}
}

func (FR *UserRepository) Update(U ents.User) error {
	_, err := FR.DB.Exec(`update User_tab set 
	name = $1, 
	password = $2, 
	cur_Foudrising_Amount = $3, 
	closed_Foundrising_Amount = $4,
	fund_balance = $5, 
	income_history = $6, 
	outcome_history = $7,
	volunteer_amount = $8,
	country = $9, 
	login = $10
	where id = $11`, U.Login,
		U.Password,
		U.Balance,
		U.CharitySum,
		U.ClosedFingAmount,
		U.Id)
	if err != nil {
		return fmt.Errorf("error in Update User_tab repo")
	} else {
		return nil
	}
}

func (FR *UserRepository) Select() ([]ents.User, bool) {
	var User_tab []ents.User
	//result := FR.DB.Table("User_tab").Order("ID").Find(&User_tab)
	rows, err := FR.DB.Query("select * from User_tab")
	if err != nil {
		return nil, false
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
			return nil, false
		}
		User_tab = append(User_tab, U)
	}

	return User_tab, (len(User_tab) != 0)
}
func (FR *UserRepository) SelectById(id uint64) (ents.User, bool) {
	var U ents.User
	row := FR.DB.QueryRow("select * from User_tab where id = $1", id)
	err := row.Scan(&U.Id,
		&U.Login,
		&U.Password,
		&U.Balance,
		&U.CharitySum,
		&U.ClosedFingAmount)
	if err != nil {
		return ents.User{}, false
	}
	return U, true
}

func (FR *UserRepository) SelectByLogin(name string) (ents.User, bool) {
	var U ents.User
	row := FR.DB.QueryRow("select * from User_tab where login = $1", name)
	err := row.Scan(&U.Id,
		&U.Login,
		&U.Password,
		&U.Balance,
		&U.CharitySum,
		&U.ClosedFingAmount)
	if err != nil {
		return ents.User{}, false
	}
	return U, true
}
