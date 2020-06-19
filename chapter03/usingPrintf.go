package main

import "fmt"

func main() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 45)
	fmt.Printf("A string: %s\n", "Bounce")
	fmt.Printf("A boolean: %t\n", true)
	fmt.Printf("Values: %v %v %v\n", 2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", false)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Types as code: %#v %#v %#v\n", "\t", "", "\n")
	fmt.Printf("Percent sign: %%\n")
}
