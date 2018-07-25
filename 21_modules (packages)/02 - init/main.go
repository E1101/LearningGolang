package main

import (
	"database/sql"
	// import and just initialize postgres package
	// here, it will register a driver into sql package that can used later
	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

func main() {
	// We call the Open method provided by
	// the sql package; this works because
	// the driver registered itself with the sql
	// package in its init function.
	sql.Open("postgres", "mydb")
}
