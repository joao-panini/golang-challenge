package model

import (
	"api/src/security"
	"errors"
	"strings"
	"time"
)

type Account struct {
	ID         uint64    `json:"Id,omitempty"`
	Name       string    `json:"Name,omitempty"`
	Cpf        string    `json:"Cpf,omitempty"`
	Secret     string    `json:"Secret,omitempty"`
	Balance    float64   `json:"Balance,omitempty"`
	Created_at time.Time `json:"Created_at,omitempty"`
}

func (account *Account) validate(stage string) error {

	if account.Name == "" {
		return errors.New("Nome é obrigatorio")
	}
	if account.Cpf == "" {
		return errors.New("Cpf é obrigatorio")
	}
	if account.Secret == "" && stage == "cadastro" {
		return errors.New("Secret é obrigatorio")
	}
	return nil
}

func (account *Account) format(stage string) error {
	account.Name = strings.TrimSpace(account.Name)

	if stage == "cadastro" {
		secretHash, erro := security.Hash(account.Secret)
		if erro != nil {
			return erro
		}
		account.Secret = string(secretHash)
	}
	return nil
}

func (account *Account) Prepare(stage string) error {
	if erro := account.validate(stage); erro != nil {
		return erro
	}

	if erro := account.format(stage); erro != nil {
		return erro
	}
	return nil
}
