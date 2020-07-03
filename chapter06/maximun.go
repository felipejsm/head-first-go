package main

import (
	"fmt"
	"math"
)

func maximun(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
func main() {
	fmt.Println(maximun(71.8, 22.3, 25.6))
	fmt.Println(maximun(-99.2, 21.0, -500.1, -1.1))
}
