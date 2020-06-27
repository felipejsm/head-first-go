package main

import (
	"dates"
	"fmt"
)

// Main function that prints dates
func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
