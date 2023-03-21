package db

import (
	"database/sql"
)

type OperationType struct {
	ID          int64        `json:"id"`
	Multiplier  int32        `json:"multiplier"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
