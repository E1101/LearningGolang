package main

// To install the package mgo , run the following command:
// $ go get gopkg.in/mgo.v2

import (
	"gopkg.in/mgo.v2"
	// If you want to use the bson package
	_ "gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)

func main()  {
	/*session, err :=*/ mgo.Dial("localhost")

	// connection with a cluster of servers
	/*session, err :=*/ mgo.Dial("server1.mongolab.com,server2.mongolab.com")

	// establish connection to one or a cluster of servers,
	// allows you to pass customized information to the server
	mongoDialInfo := &mgo.DialInfo{
		Addrs: []string{"localhost"},
		Timeout: 60 * time.Second,
		Database: "bookmarkdb",
		Username: "shijuvar",
		Password: "password123",
	}

	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		panic("Error Connect DB.")
	}

	// !! All session methods are concurrency-safe so that you can call them from multiple goroutines.


	collection := session.DB("bookmarkdb").C("bookmarks")

	fmt.Println(collection)
}
