package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("C:\\moses-learning\\head-first-go\\chapter07\\votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
