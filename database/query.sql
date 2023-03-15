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
