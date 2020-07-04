package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Joseph": 7.6, "Stalin": 9.2, "Krotsky": 4.5}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f\n", name, grades[name])
	}
}
