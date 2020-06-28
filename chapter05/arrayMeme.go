package main

import "fmt"

func main() {
	var dorime [3]string = [3]string{"do", "re", "me"}
	var ameno [3]string
	ameno[0] = "a"
	ameno[1] = "me"
	ameno[2] = "no"
	for i := 0; i < 3; i++ {
		for index := 0; index < 3; index++ {
			fmt.Print(ameno[index])
		}
		println("")
	}
	for index := 0; index < 3; index++ {
		fmt.Print(dorime[index])
	}

}
