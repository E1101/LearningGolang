package main

import (
	"fmt"
	"time"
)

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	// The formatting is done
	// with use of reference value
	// Jan 2 15:04:05 2006 MST
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	// > tTime is: 2017/3/5


	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))
	// > The time is: 08:05


	//The predefined formats could
	// be used
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))
	// > The time is: Sun, 05 Mar 2017 08:05:02 +0330


	// The formatting supports space padding
	//only for days in Go version 1.9.2
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))
	// > tTime is: 2017/3/ 5


	// The zero padding is done by adding 0
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))
	// > tTime is: 2017/03/05


	//The fraction with leading zeros use 0s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.00"))
	// > tTime is: 08:05:02.00


	//The fraction without leading zeros use 9s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))
	// > tTime is: 08:05:02


	// Append format appends the formatted time to given
	// buffer
	fmt.Println(string( tTime.AppendFormat([]byte("The time is up: "), "03:04PM") ))
	// > The time is up: 08:05AM
}
