package main

import "fmt"

func main() {
	logicBounceLyrics := [13]string{
		"Out the blue, like a Crip",
		"But I never bang though",
		"They know, we've been at it way before the fame, ho",
		"Way before the fame, ho",
		"At it way before the, way before the fame ho (Bounce)",
		"Way before the, way before the fame ho (Bounce)",
		"Way before the fame (Bounce)",
		"Way before the fame (Bounce)",
		"Way before the, at it way before the fame (Bounce)",
		"Way before the fame, at it way before the fame",
		"We've been at it way before the fame",
		"(Ridin' with my homies that be down with me)",
		"(Pop bottles with my homies that be down with me)",
	}
	// for index, chorus := range logicBounceLyrics
	for _, chorus := range logicBounceLyrics {
		fmt.Println(chorus)
	}
}
