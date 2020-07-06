package main

import "github.com/headfirstgo/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func main() {

	mixtape := []string{"Keanu Reeves", "Everybody dies", "Bounce", "Nikki"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
