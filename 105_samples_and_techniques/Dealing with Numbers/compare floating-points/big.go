package main

import (
	"fmt"
	"math/big"
)


var da1   float64 = 0.299999992
var db1   float64 = 0.299999991

var prec  uint    = 32
var prec2 uint    = 16


func main() {

	fmt.Printf("Comparing float64 with '==' equals: %v\n", da1 == db1)

	daB := big.NewFloat(da1).SetPrec(prec)
	dbB := big.NewFloat(db1).SetPrec(prec)

	fmt.Printf("A: %v \n", daB)
	fmt.Printf("B: %v \n", dbB)
	fmt.Printf("Comparing big.Float with precision: %d : %v\n",
		prec, daB.Cmp(dbB) == 0)


	daB = big.NewFloat(da1).SetPrec(prec2)
	dbB = big.NewFloat(db1).SetPrec(prec2)

	fmt.Printf("A: %v \n", daB)
	fmt.Printf("B: %v \n", dbB)
	fmt.Printf("Comparing big.Float with precision: %d : %v\n",
		prec2, daB.Cmp(dbB) == 0)
}
