package main

import "github.com/lethang7794/head-first-go/gadget"

type Player interface { // Define an interface type.
	Play(song string) // Require a Play method with string parameter.
	Stop()            // Also require a Stop method
}

func playList(device Player, songs []string) { // Accept any Player, not just a TapePlayer
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	var player Player = gadget.TapePlayer{}
	playList(player, mixtape) // Pass a TapePlayer to playList

	player = gadget.TapeRecorder{}
	playList(player, mixtape) // Pass a TapeRecorder to playList
}
