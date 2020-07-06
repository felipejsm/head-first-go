package main

import "github.com/headfirstgo/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Keanu Reeves", "Everybody dies", "Bounce", "Nikki"}
	playList(player, mixtape)
}
