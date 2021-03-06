package main

import "fmt"

// How to tell zero values apart from assigned values
func main() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	// comma ok idiom -> used to distinguish from a 'found' and 'not found' thing
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	delete(counters, "a")
	value, ok = counters["a"]
	fmt.Println(value, ok)
}
