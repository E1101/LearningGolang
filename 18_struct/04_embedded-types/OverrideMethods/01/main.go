package main

import (
	"time"
	"fmt"
)

func main() {

	d := Developer{
		Employee{
			"Steve",
			"John",
			"Software Engineer",
			"San Francisco",
			time.Date(1990, time.February, 17, 0, 0, 0, 0, time.UTC),
		},
		[]string{"Go", "Docker", "Kubernetes"},
	}

	d.PrintName()
	d.PrintDetails()

}

// interfaces:

type TeamMember interface {
	PrintName()
	PrintDetails()
}


// Structs:

type Employee struct {
	FirstName, LastName string
	JobTitle, Location  string
	Dob                 time.Time
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle,
		e.Location)
}

// Inherit:

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

// Overrides the PrintDetails
func (d Developer) PrintDetails() {
	// Call Employee PrintDetails
	d.Employee.PrintDetails()

	fmt.Println("Technical Skills:")
	for _, v := range d.Skills {
		fmt.Println(v)
	}
}