package postgres

import (
	"database/sql"
	"fmt"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type FoundrisingRepository struct {
	DB *sql.DB
}

func NewFoundrisingRepository(db_ *sql.DB) *FoundrisingRepository {
	return &FoundrisingRepository{DB: db_}
}
func (FR *FoundrisingRepository) Insert(F ents.Foundrising) (ents.Foundrising, error) {
	var LastId uint64
	row := FR.DB.QueryRow("SELECT id from Foundrising_tab ORDER BY id DESC LIMIT 1")
	err := row.Scan(&LastId)
	if err != nil {
		if err != sql.ErrNoRows {
			return ents.Foundrising{}, err
		}
	}
	_, err = FR.DB.Exec(`insert into Foundrising_tab(
		id, 
		found_id, 
		description, 
		required_sum, 
		current_sum,
		philantrops_amount, 
		creation_date, 
		closing_date) 
		values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		LastId+1,
		F.Found_id,
		F.Description,
		F.Required_sum,
		F.Current_sum,
		F.Philantrops_amount,
		F.Creation_date,
		F.Closing_date,
	)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf("error in insert Foundrising repo")
	} else {
		return F, nil
	}
}

func (FR *FoundrisingRepository) Delete(F ents.Foundrising) (ents.Foundrising, error) {
	_, err := FR.DB.Exec("delete from Foundrising_tab where id = $1", F.Id)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf("error in Delete Foundrising repo")
	} else {
		return F, nil
	}
}

func (FR *FoundrisingRepository) Update(F ents.Foundrising) (ents.Foundrising, error) {
	_, err := FR.DB.Exec(`update Foundrising_tab set 
	found_id = $1, 
	description = $2, 
	required_sum = $3, 
	current_sum = $4,
	philantrops_amount = $5, 
	creation_date = $6, 
	closing_date = $7
	where id = $8`, F.Found_id,
		F.Description,
		F.Required_sum,
		F.Current_sum,
		F.Philantrops_amount,
		F.Creation_date,
		F.Closing_date,
		F.Id)
	if err != nil {
		return ents.Foundrising{}, fmt.Errorf("error in Update Foundrising_tab repo")
	} else {
		return F, nil
	}
}

func (FR *FoundrisingRepository) Select() ([]ents.Foundrising, error) {
	var Foundrising_tab []ents.Foundrising
	rows, err := FR.DB.Query("select * from Foundrising_tab")
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundrising{}
		err := rows.Scan(&F.Id,
			&F.Found_id,
			&F.Description,
			&F.Required_sum,
			&F.Current_sum,
			&F.Philantrops_amount,
			&F.Creation_date,
			&F.Closing_date,
		)
		if err != nil {
			return nil, err
		}
		Foundrising_tab = append(Foundrising_tab, F)
	}

	return Foundrising_tab, nil
}
func (FR *FoundrisingRepository) SelectById(id uint64) (ents.Foundrising, error) {
	var F ents.Foundrising
	row := FR.DB.QueryRow("select * from Foundrising_tab where id = $1", id)
	err := row.Scan(&F.Id,
		&F.Found_id,
		&F.Description,
		&F.Required_sum,
		&F.Current_sum,
		&F.Philantrops_amount,
		&F.Creation_date,
		&F.Closing_date)
	if err != nil {
		return ents.Foundrising{}, err
	}
	return F, nil
}
func (FR *FoundrisingRepository) SelectByFoundId(id uint64) ([]ents.Foundrising, error) {
	var Foundrising_tab []ents.Foundrising
	rows, err := FR.DB.Query("select * from Foundrising_tab where found_id = $1", id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundrising{}
		err := rows.Scan(&F.Id,
			&F.Found_id,
			&F.Description,
			&F.Required_sum,
			&F.Current_sum,
			&F.Philantrops_amount,
			&F.Creation_date,
			&F.Closing_date,
		)
		if err != nil {
			return nil, err
		}
		Foundrising_tab = append(Foundrising_tab, F)
	}

	return Foundrising_tab, nil
}
func (FR *FoundrisingRepository) SelectByFoundIdActive(id uint64) ([]ents.Foundrising, error) {
	var Foundrising_tab []ents.Foundrising
	rows, err := FR.DB.Query("select * from Foundrising_tab where found_id = $1 and closing_date = $2", id, "")
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundrising{}
		err := rows.Scan(&F.Id,
			&F.Found_id,
			&F.Description,
			&F.Required_sum,
			&F.Current_sum,
			&F.Philantrops_amount,
			&F.Creation_date,
			&F.Closing_date,
		)
		if err != nil {
			return nil, err
		}
		Foundrising_tab = append(Foundrising_tab, F)
	}

	return Foundrising_tab, nil
}

func (FR *FoundrisingRepository) SelectByCreateDate(date string) ([]ents.Foundrising, error) {
	var Foundrising_tab []ents.Foundrising
	rows, err := FR.DB.Query("select * from Foundrising_tab where creation_date = $1", date)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundrising{}
		err := rows.Scan(&F.Id,
			&F.Found_id,
			&F.Description,
			&F.Required_sum,
			&F.Current_sum,
			&F.Philantrops_amount,
			&F.Creation_date,
			&F.Closing_date,
		)
		if err != nil {
			return nil, err
		}
		Foundrising_tab = append(Foundrising_tab, F)
	}

	return Foundrising_tab, nil
}

func (FR *FoundrisingRepository) SelectByCloseDate(date string) ([]ents.Foundrising, error) {
	var Foundrising_tab []ents.Foundrising
	rows, err := FR.DB.Query("select * from Foundrising_tab where closing_date = $1", date)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		F := ents.Foundrising{}
		err := rows.Scan(&F.Id,
			&F.Found_id,
			&F.Description,
			&F.Required_sum,
			&F.Current_sum,
			&F.Philantrops_amount,
			&F.Creation_date,
			&F.Closing_date,
		)
		if err != nil {
			return nil, err
		}
		Foundrising_tab = append(Foundrising_tab, F)
	}

	return Foundrising_tab, nil
}

func (FR *FoundrisingRepository) SelectByIdAndFoundId(id_ uint64, found_id_ uint64) (ents.Foundrising, error) {
	var F ents.Foundrising
	row := FR.DB.QueryRow("select * from Foundrising_tab where id = $1 and found_id = $2", id_, found_id_)
	err := row.Scan(&F.Id,
		&F.Found_id,
		&F.Description,
		&F.Required_sum,
		&F.Current_sum,
		&F.Philantrops_amount,
		&F.Creation_date,
		&F.Closing_date)
	if err != nil {
		return ents.Foundrising{}, err
	}
	return F, nil
}
