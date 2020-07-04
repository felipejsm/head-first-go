package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	a := magazine.Address{Street: "Baker Street", State: "Marylebone", PostalCode: "221B", City: "London"}
	var s magazine.Subscriber
	s.Rate = 4.99
	s.Address = a
	fmt.Println(s.City)
	fmt.Println(s.Address)

	var e magazine.Employee
	e.Name = "Thomas Shelby"
	e.Salary = 60000
	fmt.Println(e.Name)
	fmt.Println(e.Salary)
}
