-- name: CreateAccount :one
INSERT INTO accounts (
  document_number
) VALUES (
  $1
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: CreateTransactions :one
INSERT INTO transactions (
  account_id, operation_type_id, amount 
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransactions :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: DropAccount :exec
DELETE FROM accounts 
WHERE accounts.id = $1;

-- name: GetOperationTypes :one
SELECT * FROM operation_types
WHERE id = $1 LIMIT 1;

-- name: DropTransaction :exec
DELETE FROM transactions 
WHERE transactions.id = $1;

-- name: CreateOperationType :one
INSERT INTO operation_types (
  multiplier, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DropOperationType :exec
DELETE FROM operation_types 
WHERE operation_types.id = $1;

