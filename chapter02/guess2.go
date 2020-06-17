package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Random number between 1 and 100")
	fmt.Println("Can you guess it?")
	fmt.Println(target)
}
