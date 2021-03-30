package repositories

import (
	"api/src/model"
	"database/sql"
)

type AccountsRepository interface {
	Save(model.Account) (uint64, error)
	FindAll([]model.Account, error)
	FindBalanceById(ID uint64) (model.Account, error)
	FindByCPF(cpf string) (model.Account, error)
	FindByID(ID uint64) (model.Account, error)
	UpdateAccount(ID uint64, acct *model.Account) error
	UpdateBalance(ID uint64, amount float64) error
}

type accountsRepoImpl struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accountsRepoImpl {
	return &accountsRepoImpl{db}
}

func (r accountsRepoImpl) Save(account model.Account) (uint64, error) {
	statement, erro := r.db.Prepare(
		"insert into accounts (Name,Cpf,Secret) values (?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(account.Name, account.Cpf, account.Secret)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository accountsRepoImpl) FindAll() ([]model.Account, error) {
	rows, erro := repository.db.Query(
		"select id,name,cpf,balance,Created_at from accounts",
	)

	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var accountsList []model.Account
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

		accountsList = append(accountsList, account)
	}
	return accountsList, nil
}

func (repository accountsRepoImpl) FindBalanceById(ID uint64) (model.Account, error) {
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

func (repository accountsRepoImpl) FindByCPF(cpf string) (model.Account, error) {
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

func (repository accountsRepoImpl) FindByID(ID uint64) (model.Account, error) {
	row, erro := repository.db.Query("select id,name,cpf,balance from accounts where id = ?", ID)
	if erro != nil {
		return model.Account{}, erro
	}
	defer row.Close()

	var account model.Account
	if row.Next() {
		if erro = row.Scan(
			&account.ID,
			&account.Name,
			&account.Cpf,
			&account.Balance,
		); erro != nil {
			return model.Account{}, erro
		}
	}
	return account, nil
}

func (repository accountsRepoImpl) UpdateAccount(ID uint64, account *model.Account) error {
	statement, erro := repository.db.Prepare(
		"update accounts set name = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(account.Name, ID); erro != nil {
		return erro
	}
	return nil
}

func (repository accountsRepoImpl) UpdateAccountBalance(ID uint64, amount float64) error {
	statement, erro := repository.db.Prepare(
		"update accounts set balance = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(amount, ID); erro != nil {
		return erro
	}
	return nil
}
