package main

import "fmt"

func main() {
	slice := []string{"Bounce", "bounce", "bounce", " Jay-Z huh?"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "Yeah", "yeah", "yeah", "Roc-A-Fella y'all", "ha ha")
	fmt.Println(slice, len(slice))
}
