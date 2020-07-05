package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Day())
}
