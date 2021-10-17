package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/daryalazareva/microapp/config"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

//create db connection and use it to create the queries object
//An entry point to all unit tests inside one specific golang package
func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load configuration: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testDB.SetMaxOpenConns(1000)
	testDB.SetMaxIdleConns(5)
	testDB.SetConnMaxLifetime(time.Hour)

	testQueries = New(testDB)

	os.Exit(m.Run())
}
