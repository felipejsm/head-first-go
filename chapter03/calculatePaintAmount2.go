package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var width, height float64

	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Insert width value: ")
	input, err := scanner.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	width, _ = strconv.ParseFloat(input, 64)

	fmt.Printf("\nInsert height value: ")
	input, err = scanner.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	height, _ = strconv.ParseFloat(input, 64)

	amount, newError := paintNeededWithMultipleReturn(width, height)
	fmt.Println(newError)
	fmt.Println("%0.2f liters needed\n", amount)
}

func paintNeededWithMultipleReturn(width, height float64) (amount float64, err error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	amount = area / 10
	err = nil
	return //(area / 10), nil
}
