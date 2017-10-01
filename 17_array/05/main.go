package main

import "fmt"

func main() {
	x := [5]float64{
		0: 98,
		2: 93,
		77,
		82,
	}

	fmt.Println(len(x))
	fmt.Println(x[0])

	for i := 0; i < 256; i++ {
		x[i] = float64(i)
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
