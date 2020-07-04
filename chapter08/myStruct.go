package main

import "fmt"

func main() {
	var myStructure struct {
		name string
		age  int
	}
	myStructure.name = "Jean Paul Sartre"
	myStructure.age = 64
	fmt.Println(myStructure)
}
