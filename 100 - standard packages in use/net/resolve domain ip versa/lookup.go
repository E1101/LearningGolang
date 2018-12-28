package main

import (
	"fmt"
	"net"
)

func main() {

	// Resolve by IP
	//
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	// admin007.digifilm.ir.lc.
	// digifilm.ir.lc.
	// sim.digifilm.ir.lc.
	// dl3.digiFilm.ir.lc.
	// dl3-1m.digifilm.ir.lc.
	// >
	for _, addr := range addrs {
		fmt.Println(addr)
	}


	// Resolve by address
	//
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	// ::1
	// 127.0.0.1
	// >
	for _, ip := range ips {
		fmt.Println(ip.String())
	}

}
