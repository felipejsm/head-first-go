package main

import "fmt"

func main() {
	isFalse := true
	if isFalse {
		fmt.Println("False is", isFalse)
	}
	fmt.Println("2 > 2 is", (2 > 2))
	fmt.Println("\tNumber\tOperator\tNumber\tResult")
	for index := 1; index < 11; index++ {
		fmt.Printf("\t%v\t%v\t\t%v\t%v", index, "x", index+1, index*(index+1))
		fmt.Println("")
	}
}
