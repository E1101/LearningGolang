package main

import (
	"fmt"
	"net"
)

func main() {

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Resolve addresses for each interface
		// lists the available network addresses.
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}


		fmt.Println(interf.Name)

		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}

}
