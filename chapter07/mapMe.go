package main

import "fmt"

func main() {
	var rappers map[string]int
	rappers = make(map[string]int)
	rappers["Jay-z"] = 50
	rappers["Mos def"] = 46
	rappers["Kanye west"] = 43
	for rapper, _ := range rappers {
		fmt.Printf("%s is %d old\n", rapper, rappers[rapper])
	}
	emos := make(map[string]int)
	emos["Jeremy Bolm"] = 40
	emos["Jesse Lacey"] = 42
	emos["Mike Kinsella"] = 43

	for emo, _ := range emos {
		fmt.Printf("%s is %d old\n", emo, emos[emo])
	}

	danceGavinDanceVocalists := map[string]int{"Kurt Travis": 2010, "Jonny Craig": 2012, "Tilian Pearson": 2020}

	for dgd, _ := range danceGavinDanceVocalists {
		fmt.Printf("%s left DGD at %d\n", dgd, danceGavinDanceVocalists[dgd])
	}
}
