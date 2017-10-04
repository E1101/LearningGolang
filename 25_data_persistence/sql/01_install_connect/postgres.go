package main

// run the following command:
// $ go get github.com/lib/pq

import (
	"database/sql"
	//you use package pq only for invoking its init function for registering
	// your driver ( "postgres" ) with database/sql.
	// Because the package pq is importing just for invoking its init function,
	// a blank identifier (_) is used as the package alias to avoid compilation error.
	_ "github.com/lib/pq"

	"log"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@localhost/dbname")
	if err != nil {
		log.Fatal(err)
	}
}



