package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("logic_bounce.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\tTamanho do Arquivo: %d\tNome do Arquivo: %s", fileInfo.Size, fileInfo.Name())
}
