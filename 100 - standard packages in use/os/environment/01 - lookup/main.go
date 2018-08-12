package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	key := "DB_CONN"

	connStr, isFound := os.LookupEnv(key)
	if !isFound {
		log.Printf("The env variable %s is not set.\n", key)
	}


	fmt.Println(connStr)
}

// $ unset DB_CONN && go run lookup.go
//
