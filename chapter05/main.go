package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("C:\\moses-learning\\head-first-go\\chapter05\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sample := float64(len(numbers))
	fmt.Println("Sum: ", sum)
	fmt.Printf("Average: %0.2f\n", sum/sample)
}
