package model

import (
	"errors"
	"time"
)

type Transfers struct {
	ID            uint64    `json:"Id,omitempty"`
	FromAccountID uint64    `json:"FromAccountID,omitempty"`
	ToAccountID   uint64    `json:"ToAccountID,omitempty"`
	Amount        float64   `json:"Amount,omitempty"`
	Created_at    time.Time `json:"Created_at,omitempty"`
}

func (transfer *Transfers) Validate() error {

	if transfer.FromAccountID <= 0 {
		return errors.New("destinatario é obrigatorio")
	}
	if transfer.Amount <= 0 {
		return errors.New("insira um valor de transferencia maior do que 0")
	}
	return nil
}
