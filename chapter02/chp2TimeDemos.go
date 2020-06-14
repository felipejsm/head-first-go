package main

import (
	"fmt"
	"time"
)

func main() {
	jayZBirthDate := time.Date(1969, 12, 4, 12, 3, 3, 2, time.UTC)

	fmt.Println("Jay has ", time.Since(jayZBirthDate))
	fmt.Println("Now sub Jay-z's birthdate", time.Now().Sub(jayZBirthDate).)
}
