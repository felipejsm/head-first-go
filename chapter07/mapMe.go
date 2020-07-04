package main

import "fmt"

func main() {
	var rappers map[string]int
	rappers = make(map[string]int)
	rappers["Jay-z"] = 50
	rappers["Mos def"] = 46
	rappers["Kanye west"] = 43
	for rapper, _ := range rappers {
		fmt.Printf("%s is %d old\n", rapper, rappers[rapper])
	}
}
