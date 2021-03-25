package repositories

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type accounts struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accounts {
	return &accounts{db}
}

func (repository accounts) Create(account model.Account) (uint64, error) {
	fmt.Println("estou aqui")
	statement, erro := repository.db.Prepare(
		"insert into accounts (Name,Cpf,Secret,Balance) values (?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(account.Name, account.Cpf, account.Secret, account.Balance)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository accounts) GetAllAccounts() ([]model.Account, error) {
	rows, erro := repository.db.Query(
		"select id,name,cpf,balance,Created_at from accounts",
	)

	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var account model.Account

		if erro = rows.Scan(
			&account.ID,
			&account.Name,
			&account.Cpf,
			&account.Balance,
			&account.Created_at,
		); erro != nil {
			return nil, erro
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (repository accounts) GetAccountBalanceById(ID uint64) (model.Account, error) {
	rows, erro := repository.db.Query(
		"select balance from accounts where id = ?",
		ID,
	)
	if erro != nil {
		return model.Account{}, erro
	}
	defer rows.Close()

	var account model.Account
	if rows.Next() {
		if erro = rows.Scan(
			&account.Balance,
		); erro != nil {
			return model.Account{}, erro
		}
	}
	return account, nil
}

func (repository accounts) GetAccountByCPF(cpf string) (model.Account, error) {
	row, erro := repository.db.Query("select id, secret from accounts where cpf = ?", cpf)
	if erro != nil {
		return model.Account{}, erro
	}
	defer row.Close()

	var account model.Account
	if row.Next() {
		if erro = row.Scan(
			&account.ID,
			&account.Secret,
		); erro != nil {
			return model.Account{}, erro
		}
	}
	return account, nil
}
