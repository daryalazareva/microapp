package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:467958@localhost:5432/microapp?sslmode=disable"
)

//create db connection and use it to create the queries object
//An entry point to all unit tests inside one specific golang package
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testDB.SetMaxOpenConns(100)
	testDB.SetMaxIdleConns(5)
	testDB.SetConnMaxLifetime(time.Hour)

	testQueries = New(testDB)

	os.Exit(m.Run())
}
