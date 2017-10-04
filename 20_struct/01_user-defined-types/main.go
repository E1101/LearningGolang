package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// ..

// Person Struct with a Slice of a User-Defined Type as the Type for Field
type Address struct {
	Street, City, State, Zip string
	IsShippingAddress bool
}

type Customer struct {
	Person
	Addresses []Address
}

func main()  {
	// Creating Instances of Struct Types
	var c Person
	c.FirstName = "Alex"
	c.LastName = "John"
	c.Email = "alex@email.com"
	c.Phone = "732-757-2923"

	// Creating a Struct Instance Using a Struct Literal
	p := Person{
		FirstName: "Alex",
		LastName: "John",
		Email: "alex@email.com",
		Phone: "732-757-2923",
	}

	fmt.Println(p.FirstName+" "+p.LastName)


	// ...

	cst := Customer{
		Person: Person{
			FirstName: "Alex",
			LastName: "John",
			Email: "alex@email.com",
			Phone: "732-757-2923",
		},
		Addresses: []Address{
			{
				Street: "1 Mission Street",
				City: "San Francisco",
				State: "CA",
				Zip: "94105",
				IsShippingAddress: true,
			},
			Address{
				Street: "49 Stevenson Street",
				City: "San Francisco",
				State: "CA",
				Zip: "94105",
			},
		},
	}

	fmt.Println(cst.FirstName+" "+cst.LastName)

}