// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: query.sql

package database

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  document_number
) VALUES (
  $1
)
RETURNING id, document_number, created_at, updated_at
`

func (q *Queries) CreateAccount(ctx context.Context, documentNumber string) (Account, error) {
	row := q.queryRow(ctx, q.createAccountStmt, createAccount, documentNumber)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createOperationType = `-- name: CreateOperationType :one
INSERT INTO operation_types (
  multiplier, description
) VALUES (
  $1, $2
)
RETURNING id, multiplier, description, created_at, updated_at
`

type CreateOperationTypeParams struct {
	Multiplier  int32  `json:"multiplier"`
	Description string `json:"description"`
}

func (q *Queries) CreateOperationType(ctx context.Context, arg CreateOperationTypeParams) (OperationType, error) {
	row := q.queryRow(ctx, q.createOperationTypeStmt, createOperationType, arg.Multiplier, arg.Description)
	var i OperationType
	err := row.Scan(
		&i.ID,
		&i.Multiplier,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTransactions = `-- name: CreateTransactions :one
INSERT INTO transactions (
  account_id, operation_type_id, amount 
) VALUES (
  $1, $2, $3
)
RETURNING id, account_id, operation_type_id, amount, event_date, created_at, updated_at
`

type CreateTransactionsParams struct {
	AccountID       int32   `json:"account_id"`
	OperationTypeID int32   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (q *Queries) CreateTransactions(ctx context.Context, arg CreateTransactionsParams) (Transaction, error) {
	row := q.queryRow(ctx, q.createTransactionsStmt, createTransactions, arg.AccountID, arg.OperationTypeID, arg.Amount)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.OperationTypeID,
		&i.Amount,
		&i.EventDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const dropAccount = `-- name: DropAccount :exec
DELETE FROM accounts 
WHERE accounts.id = $1
`

func (q *Queries) DropAccount(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.dropAccountStmt, dropAccount, id)
	return err
}

const dropOperationType = `-- name: DropOperationType :exec
DELETE FROM operation_types 
WHERE operation_types.id = $1
`

func (q *Queries) DropOperationType(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.dropOperationTypeStmt, dropOperationType, id)
	return err
}

const dropTransaction = `-- name: DropTransaction :exec
DELETE FROM transactions 
WHERE transactions.id = $1
`

func (q *Queries) DropTransaction(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.dropTransactionStmt, dropTransaction, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, document_number, created_at, updated_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.queryRow(ctx, q.getAccountStmt, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOperationTypes = `-- name: GetOperationTypes :one
SELECT id, multiplier, description, created_at, updated_at FROM operation_types
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOperationTypes(ctx context.Context, id int64) (OperationType, error) {
	row := q.queryRow(ctx, q.getOperationTypesStmt, getOperationTypes, id)
	var i OperationType
	err := row.Scan(
		&i.ID,
		&i.Multiplier,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTransactions = `-- name: GetTransactions :one
SELECT id, account_id, operation_type_id, amount, event_date, created_at, updated_at FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransactions(ctx context.Context, id int64) (Transaction, error) {
	row := q.queryRow(ctx, q.getTransactionsStmt, getTransactions, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.OperationTypeID,
		&i.Amount,
		&i.EventDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
