package main

import "fmt"

func main() {
	amount := 45
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
