package main

import "fmt"

var metersPerLiter float64

func main() {
	var width, height float64
	metersPerLiter = 10
	width = 4.2
	height = 3.0
	fmt.Printf("%.2f liters needed\n", paintNeeded(width, height))
	width = 5.2
	height = 3.5
	paintNeededWithPrintf(width, height)
}

func paintNeededWithPrintf(width, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/metersPerLiter)
}

func paintNeeded(width, height float64) float64 {
	return width * height
}
