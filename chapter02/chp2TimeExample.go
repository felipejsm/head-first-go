package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println("Year: ", year)
	fmt.Println("Now: ", now)
}
