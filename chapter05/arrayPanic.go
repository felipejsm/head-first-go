package main

import "fmt"

func main() {
	omgPanic := [3]string{"one", "two", "three"}
	for index := 0; index < 4; index++ {
		fmt.Println(omgPanic[index])
	}
}
