package model

import (
	"time"
)

type Transfers struct {
	ID            uint64    `json:"Id,omitempty"`
	FromAccountID uint64    `json:"FromAccountID,omitempty"`
	ToAccountID   uint64    `json:"ToAccountID,omitempty"`
	Amount        float64   `json:"Amount,omitempty"`
	Created_at    time.Time `json:"Created_at,omitempty"`
}
