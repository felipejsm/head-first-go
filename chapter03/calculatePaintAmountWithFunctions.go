package main

import "fmt"

func main() {
	var width, height float64
	width = 4.2
	height = 3.0
	paintNeeded(width, height)
	width = 5.2
	height = 3.5
	paintNeeded(width, height)
}

func paintNeeded(width, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}
