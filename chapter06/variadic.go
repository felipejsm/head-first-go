package main

import "fmt"

func main() {
	severalInts(1, 3, 4, 6, 4545, 3)
}
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
