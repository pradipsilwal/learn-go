package main

import (
	"gadget"
)

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

func tryOut(player Player, song string) {
	player.Play(song)
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func tapeInterface() {
	var player Player = gadget.TapeRecorder{}
	tryOut(player, "Shake it off")
	player = gadget.TapePlayer{}
	tryOut(player, "Delicate")
}
