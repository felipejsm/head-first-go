package main

import "fmt"

func main() {
	songTitle := "che il lupo cattivo vegli su di te"
	songLyrics := [8]string{
		"Dormi piccolino, è luna piena già",
		"Strilla la sirena e l'aeroplano va",
		"Dorme per sempre l'agnello di Dio",
		"E quando suona la campana la messa inizierà che sveglia tutta la città",
		"Qui non si riposa mai",
		"La gente sparla, s'indigna e singhiozza",
		"La gente infama, tradisce e poi ride",
		"In gran tranquillità.",
	}

	fmt.Println(songTitle)
	for index := 0; index < len(songLyrics); index++ {
		fmt.Println(songLyrics[index])
	}

}
