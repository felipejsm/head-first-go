package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# R#cks!"
	lotsOfSpaces := "Origami Paper"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	anotherReplacer := strings.NewReplacer(" ", "")
	anotherFixed := anotherReplacer.Replace(lotsOfSpaces)
	fmt.Println(fixed)
	fmt.Println(anotherFixed)
}
