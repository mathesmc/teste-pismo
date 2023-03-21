package storage

import (
	"context"
	"log"

	"database/sql"

	postgres "github.com/mathesmc/teste-pismo/db/sqlc"
)

var (
	dbDriver = "postgres"
	Repo     SqlRepo
)

type Repository interface {
	ShowAccount(int64) (postgres.Account, error)
	AddAccount(createAccountParams) (postgres.Account, error)
	AddTransaction(createTransactionParams) postgres.Transaction
}

type createAccountParams struct {
	DocumentNumber string `json:"document_number,omitempty"`
}

type createTransactionParams struct {
	AccountID       int64   `json:"account_id,omitempty"`
	OperationTypeID int64   `json:"operation_type_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}

type SqlRepo struct {
	db *postgres.Queries
}

func NewSqlRepo(url string) error {
	conn, err := sql.Open(dbDriver, url)
	if err != nil {
		log.Panic("cannot connect to db because:", err)
		return err
	}

	Repo.db = postgres.New(conn)
	defer Repo.db.Close()
	return nil
}

func (s SqlRepo) GetAccount(id int64) (account postgres.Account, err error) {
	account, err = s.db.GetAccount(context.Background(), id)
	if err != nil {
		return
	}

	return
}
