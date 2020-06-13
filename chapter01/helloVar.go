package main

import "fmt"

func main() {
	var rapperName string
	var rapperNumberOfNumberOneHits int
	var wealth float64
	var rapperRealName string = "Sir Robert Bryson Hall II"
	// Default value
	var oneBadBar string

	// Gopher :=
	oneGoodBar := "Smack you with the palm, save the back for your mom"

	// Adding value
	rapperName = "Logic"
	rapperNumberOfNumberOneHits = 13
	wealth = 14.1

	fmt.Println("\tRapper\t#1 HITS\tWealth in millions")
	fmt.Printf("\t%s\t%d\t%g", rapperName, rapperNumberOfNumberOneHits, wealth)
	fmt.Printf("\n%s was born %s", rapperName, rapperRealName)
	fmt.Printf("\nOne bad bar from the rapper: %s", oneBadBar)
	fmt.Printf("\nOne good bar from the rapper: %s", oneGoodBar)
}
