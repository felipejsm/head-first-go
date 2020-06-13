package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 2

	var myFloat64Value float64 = 3.14
	var myIntValue int = 33

	myIntValue = myFloat64Value
	// Comparações e operações matemáticas devem ser feitas com o mesmo tipo
	fmt.Println("Area is ", length*width)
	fmt.Println("length > width", length > width)
}
