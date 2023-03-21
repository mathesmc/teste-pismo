package db

import (
	"database/sql"
)

type Transaction struct {
	ID              int64        `json:"id"`
	AccountID       int64        `json:"account_id"`
	OperationTypeID int64        `json:"operation_type_id"`
	Amount          float64      `json:"amount"`
	EventDate       sql.NullTime `json:"event_date"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type CreateTransactionsParams struct {
	AccountID       int64   `json:"account_id,omitempty"`
	OperationTypeID int64   `json:"operation_type_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}
