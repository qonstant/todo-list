package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

var testQueryies *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://todo-user:password@localhost:5432/todo-db?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	testQueryies = New(conn)

	os.Exit(m.Run())
}
