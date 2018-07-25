package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	// inner type
	person
	LicenseToKill bool
}

// ===========================================================================

// user defines a user in the program.
type user struct {
	name string
	email string
}
// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}
// admin represents an admin user with privileges.
// user is an inner type of the outer type admin
type admin struct {
	user // Embedded Type
	level string
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)


	// ===========================================================================

	// Create an admin user.
	ad := admin{
		user: user{
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	// the inner type never loses its identity and can always be accessed directly
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()

}
