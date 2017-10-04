package main

// run the following command:
// $ go get github.com/go-sql-driver/mysql

import (
	"database/sql"
	//you use package pq only for invoking its init function for registering
	// your driver ( "postgres" ) with database/sql.
	// Because the package pq is importing just for invoking its init function,
	// a blank identifier (_) is used as the package alias to avoid compilation error.
	_ " github.com/go-sql-driver/mysql"

	"log"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
}

