package db

import (
	"database/sql"
)

type Account struct {
	ID             int64        `json:"id"`
	DocumentNumber string       `json:"document_number"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
