package main

import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}

// $ DB_CONN="mysql:host=$servername;dbname=myDB"
// $ go run main.go
//
// $ DB_CONN=db:/user@example && go run get.go
//
// -------------------------------------------------------------------------
// 2018/08/08 11:09:46 Connection string: mysql:host=$servername;dbname=myDB
// >_
