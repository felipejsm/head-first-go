package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	primes := make([]int, 5)
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	primes[3] = 5

	fibonnaci := []int{1, 2, 3, 5, 8}

	for _, fib := range fibonnaci {
		fmt.Print(fib)
	}

	for _, note := range notes {
		fmt.Printf(note)
	}

	for _, prime := range primes {
		fmt.Print(prime)
	}
}
