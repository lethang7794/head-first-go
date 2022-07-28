package main

import "github.com/lethang7794/head-first-go/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs { // Loop over each song.
		device.Play(song) // Play the current song.
	}
	device.Stop() // Stop the player once we're done.
}

func main() {
	// Create a TapePlayer.
	player := gadget.TapePlayer{}

	// Create a slice of song titles.
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	// Play the songs using the TapePlayer
	playList(player, mixtape)

	// Will the playList function work with a TapeRecorder as well?
	//recorder := gadget.TapeRecorder{}
	//playList(recorder, mixtape) // Cannot use 'recorder' (type gadget.TapeRecorder) as the type gadget.TapePlayer
}
