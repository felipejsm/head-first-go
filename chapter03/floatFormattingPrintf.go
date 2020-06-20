package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "1234", "Cost in cents")
	floatValue := 12.3456
	fmt.Printf("%%7.3f: %7.3f\n", floatValue)
	fmt.Printf("%%7.1f: %7.1f\n", floatValue)
	fmt.Printf("%%.1f: %.1f\n", floatValue)
	fmt.Printf("%%.2f: %.2f\n", floatValue)
}
