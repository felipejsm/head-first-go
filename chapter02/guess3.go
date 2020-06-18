package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random numbere between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guess := 0; guess < 10; guess++ {
		fmt.Println("You have", 10-guess, "guesses left.")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("low")
		} else if guess > target {
			fmt.Println("high")
		} else {
			fmt.Println("Bull's eye")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was: ", target)
	}
}
