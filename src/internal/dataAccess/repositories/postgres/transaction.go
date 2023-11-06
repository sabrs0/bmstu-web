package postgres

import (
	"database/sql"
	"fmt"

	ents "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db_ *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db_}
}
func (TR *TransactionRepository) Insert(T ents.Transaction) (ents.Transaction, error) {
	var LastId uint64
	row := TR.DB.QueryRow("SELECT id from transaction_tab ORDER BY id DESC LIMIT 1")
	err := row.Scan(&LastId)
	if err != nil {
		if err != sql.ErrNoRows {
			return ents.Transaction{}, err
		}

	}
	T.Id = LastId + 1
	_, err = TR.DB.Exec(`insert into transaction_tab(
		id, 
		from_essence_type, 
		from_id, 
		to_essence_type, 
		sum_of_money,
		comment, 
		to_id) 
		values ($1, $2, $3, $4, $5, $6, $7)`,
		T.Id,
		T.From_essence_type,
		T.From_id,
		T.To_essence_type,
		T.Sum_of_money,
		T.Comment,
		T.To_id,
	)
	if err != nil {
		return ents.Transaction{}, fmt.Errorf("error in insert Transaction repo")
	}
	return T, nil
}

func (TR *TransactionRepository) Delete(T ents.Transaction) (ents.Transaction, error) {
	_, err := TR.DB.Exec("delete from transaction_tab where id = $1", T.Id)
	if err != nil {
		return ents.Transaction{}, fmt.Errorf("error in Delete transaction repo")
	}
	return T, nil
}

func (TR *TransactionRepository) Select() ([]ents.Transaction, error) {
	var transaction_tab []ents.Transaction
	rows, err := TR.DB.Query("select * from transaction_tab")
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		T := ents.Transaction{}
		err := rows.Scan(&T.Id,
			&T.From_essence_type,
			&T.From_id,
			&T.To_essence_type,
			&T.Sum_of_money,
			&T.Comment,
			&T.To_id,
		)
		if err != nil {
			return nil, err
		}
		transaction_tab = append(transaction_tab, T)
	}

	return transaction_tab, nil
}
func (TR *TransactionRepository) SelectById(id uint64) (ents.Transaction, error) {
	var T ents.Transaction
	row := TR.DB.QueryRow("select * from transaction_tab where id = $1", id)
	err := row.Scan(&T.Id,
		&T.From_essence_type,
		&T.From_id,
		&T.To_essence_type,
		&T.Sum_of_money,
		&T.Comment,
		&T.To_id)
	if err != nil {
		return ents.Transaction{}, err
	}
	return T, nil
}

func (TR *TransactionRepository) SelectFrom(type_ bool, id uint64) ([]ents.Transaction, error) {
	var transaction_tab []ents.Transaction
	rows, err := TR.DB.Query("select * from transaction_tab where from_essence_type = $1 and from_id = $2", type_, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		T := ents.Transaction{}
		err := rows.Scan(&T.Id,
			&T.From_essence_type,
			&T.From_id,
			&T.To_essence_type,
			&T.Sum_of_money,
			&T.Comment,
			&T.To_id,
		)
		if err != nil {
			return nil, err
		}
		transaction_tab = append(transaction_tab, T)
	}

	return transaction_tab, nil
}

func (TR *TransactionRepository) SelectTo(type_ bool, id uint64) ([]ents.Transaction, error) {
	var transaction_tab []ents.Transaction
	rows, err := TR.DB.Query("select * from transaction_tab where to_essence_type = $1 and to_id = $2", type_, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		T := ents.Transaction{}
		err := rows.Scan(&T.Id,
			&T.From_essence_type,
			&T.From_id,
			&T.To_essence_type,
			&T.Sum_of_money,
			&T.Comment,
			&T.To_id,
		)
		if err != nil {
			return nil, err
		}
		transaction_tab = append(transaction_tab, T)
	}

	return transaction_tab, nil
}
func (TR *TransactionRepository) SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, error) {
	var transaction_tab []ents.Transaction
	rows, err := TR.DB.Query(`select * from transaction_tab where 
							from_essence_type = $1 and 
							to_essence_type = $2 and
							to_id = $3`, ents.FROM_USER, ents.TO_FOUNDRISING, foundrising_id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		T := ents.Transaction{}
		err := rows.Scan(&T.Id,
			&T.From_essence_type,
			&T.From_id,
			&T.To_essence_type,
			&T.Sum_of_money,
			&T.Comment,
			&T.To_id,
		)
		if err != nil {
			return nil, err
		}
		transaction_tab = append(transaction_tab, T)
	}
	IDs := make([]uint64, len(transaction_tab))
	for i := range IDs {
		IDs[i] = transaction_tab[i].From_id
	}
	return IDs, nil

}
