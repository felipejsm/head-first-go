package main

import (
	"fmt"
)

func main() {
	fmt.Println("Letra\tRUNES")
	fmt.Print("A\t")
	fmt.Print('A')
	fmt.Print("\nÇ\t")
	fmt.Print('Ç')
	fmt.Println("\nHere comes the numbers: ")
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}
}
