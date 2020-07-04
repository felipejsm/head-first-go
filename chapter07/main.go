package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("C:\\moses-learning\\head-first-go\\chapter07\\votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
