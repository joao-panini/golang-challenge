package repositories

import (
	"api/src/model"
	"database/sql"
)

type transfers struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *transfers {
	return &transfers{db}
}
func (repository transfers) CreateTransaction(transfer model.Transfers) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into transfers (from_account_id,to_account_id,amount) values (?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(transfer.FromAccountID, transfer.ToAccountID, transfer.Amount)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository transfers) GetAllTransfers(ID uint64) ([]model.Transfers, error) {
	rows, erro := repository.db.Query(
		"select * from transfers where from_account_id = ?",
		ID,
	)

	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var transfersList []model.Transfers

	for rows.Next() {
		var transfer model.Transfers

		if erro = rows.Scan(
			&transfer.ID,
			&transfer.FromAccountID,
			&transfer.ToAccountID,
			&transfer.Amount,
			&transfer.Created_at,
		); erro != nil {
			return nil, erro
		}

		transfersList = append(transfersList, transfer)
	}
	return transfersList, nil
}
