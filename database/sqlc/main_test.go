package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var (
	dbDriver = "postgres"
	connStr  = "postgresql://root:secret@localhost:5432/mini_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal("cannot connect to db because:", err)
	}

	testQueries = New(db)
	defer db.Close()

	os.Exit(m.Run())
}
