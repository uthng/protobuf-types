package main

import (
	//"fmt"
	"io/ioutil"
	//"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	pb "github.com/uthng/protobuf-types/example/genpb/users"
	_ "github.com/uthng/protobuf-types/types/null"
)

func main() {
	var result *pb.User

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

	sqlStmt := "SELECT * FROM users LIMIT 1"

	err = db.Get(&result, sqlStmt)
	if err != nil {
		panic(err.Error())
	}
}
