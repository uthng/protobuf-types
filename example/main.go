package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/uthng/protobuf-types/types/null"
)

func main() {

	// Create DB pool
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=mobiz dbname=test-sqlx sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	// create table starfleet and insert some test data
	setupStmt, err := ioutil.ReadFile("example.sql")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(string(setupStmt))
	if err != nil {
		panic(err.Error())
	}

	sqlStmt := "SELECT * FROM users"

	err := db.Select(&result, sqlStmt)
	if err != nil {
		return result, err
	}
}
