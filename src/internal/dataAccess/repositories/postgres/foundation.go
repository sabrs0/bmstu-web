package postgres

import (
	"database/sql"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"

	"fmt"
)

type FoundationRepository struct {
	DB *sql.DB
}

func NewFoundationRepository(db_ *sql.DB) *FoundationRepository {
	return &FoundationRepository{DB: db_}
}
func (FR *FoundationRepository) Insert(F ents.Foundation) (ents.Foundation, error) {
	var LastId uint64
	row := FR.DB.QueryRow("SELECT id from foundation_tab ORDER BY id DESC LIMIT 1")
	err := row.Scan(&LastId)
	if err != nil {
		if err != sql.ErrNoRows {
			return ents.Foundation{}, err
		}
	}
	F.Id = LastId + 1
	_, err = FR.DB.Exec(`insert into foundation_tab(
		id, 
		name, 
		password, 
		cur_Foudrising_Amount, 
		closed_Foundrising_Amount,
		fund_balance, 
		income_history, 
		outcome_history,
		volunteer_amount,
		country, 
		login) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		F.Id,
		F.Name,
		F.Password,
		F.CurFoudrisingAmount,
		F.ClosedFoundrisingAmount,
		F.Fund_balance,
		F.Income_history,
		F.Outcome_history,
		F.Volunteer_amount,
		F.Country,
		F.Login,
	)
	if err != nil {
		return ents.Foundation{}, fmt.Errorf("error in insert Foundation repo: %s", err.Error())
	}
	return F, nil
}

func (FR *FoundationRepository) Delete(F ents.Foundation) (ents.Foundation, error) {
	_, err := FR.DB.Exec("delete from foundation_tab where id = $1", F.Id)
	if err != nil {
		return ents.Foundation{}, fmt.Errorf("error in Delete Foundation repo: %s", err.Error())
	}
	return F, nil
}

func (FR *FoundationRepository) Update(F ents.Foundation) (ents.Foundation, error) {
	_, err := FR.DB.Exec(`update foundation_tab set 
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
	where id = $11`, F.Name,
		F.Password,
		F.CurFoudrisingAmount,
		F.ClosedFoundrisingAmount,
		F.Fund_balance,
		F.Income_history,
		F.Outcome_history,
		F.Volunteer_amount,
		F.Country,
		F.Login,
		F.Id)
	if err != nil {
		return ents.Foundation{}, fmt.Errorf("error in Update foundation_tab repo: %s", err.Error())
	}
	return F, nil
}

func (FR *FoundationRepository) Select() ([]ents.Foundation, error) {
	var foundation_tab []ents.Foundation
	rows, err := FR.DB.Query("select * from foundation_tab")
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundation{}
		err := rows.Scan(&F.Id,
			&F.Name,
			&F.Password,
			&F.CurFoudrisingAmount,
			&F.ClosedFoundrisingAmount,
			&F.Fund_balance,
			&F.Income_history,
			&F.Outcome_history,
			&F.Volunteer_amount,
			&F.Country,
			&F.Login,
		)
		if err != nil {
			return nil, err
		}
		foundation_tab = append(foundation_tab, F)
	}

	return foundation_tab, nil
}
func (FR *FoundationRepository) SelectById(id uint64) (ents.Foundation, error) {
	var F ents.Foundation
	row := FR.DB.QueryRow("select * from foundation_tab where id = $1", id)
	err := row.Scan(&F.Id,
		&F.Name,
		&F.Password,
		&F.CurFoudrisingAmount,
		&F.ClosedFoundrisingAmount,
		&F.Fund_balance,
		&F.Income_history,
		&F.Outcome_history,
		&F.Volunteer_amount,
		&F.Country,
		&F.Login)
	if err != nil {
		return F, err
	}
	return F, nil
}

func (FR *FoundationRepository) SelectByLogin(name string) (ents.Foundation, error) {
	var F ents.Foundation
	row := FR.DB.QueryRow("select * from foundation_tab where login = $1", name)
	err := row.Scan(&F.Id,
		&F.Name,
		&F.Password,
		&F.CurFoudrisingAmount,
		&F.ClosedFoundrisingAmount,
		&F.Fund_balance,
		&F.Income_history,
		&F.Outcome_history,
		&F.Volunteer_amount,
		&F.Country,
		&F.Login)
	if err != nil {
		return F, err
	}
	return F, nil
}
func (FR *FoundationRepository) SelectByName(name string) (ents.Foundation, error) {
	var F ents.Foundation
	row := FR.DB.QueryRow("select * from foundation_tab where name = $1", name)
	err := row.Scan(&F.Id,
		&F.Name,
		&F.Password,
		&F.CurFoudrisingAmount,
		&F.ClosedFoundrisingAmount,
		&F.Fund_balance,
		&F.Income_history,
		&F.Outcome_history,
		&F.Volunteer_amount,
		&F.Country,
		&F.Login)
	if err != nil {
		return F, err
	}
	return F, nil
}
func (FR *FoundationRepository) SelectByCountry(country string) ([]ents.Foundation, error) {
	var foundation_tab []ents.Foundation
	rows, err := FR.DB.Query("select * from foundation_tab where country = $1", country)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundation{}
		err := rows.Scan(&F.Id,
			&F.Name,
			&F.Password,
			&F.CurFoudrisingAmount,
			&F.ClosedFoundrisingAmount,
			&F.Fund_balance,
			&F.Income_history,
			&F.Outcome_history,
			&F.Volunteer_amount,
			&F.Country,
			&F.Login,
		)
		if err != nil {
			return foundation_tab, err

		}
		foundation_tab = append(foundation_tab, F)
	}

	return foundation_tab, nil
}
