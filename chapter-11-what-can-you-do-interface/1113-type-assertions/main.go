package main

import "github.com/lethang7794/head-first-go/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("A Test Track")
	player.Stop()

	// Let's convert player to type TapeRecorder, so we can use Record method!
	//recorder := gadget.TapeRecorder{player} // It won't work.
	//recorder.Record()
}

func main() {
	var player Player = gadget.TapeRecorder{}
	TryOut(player)

	var recorder = player.(gadget.TapeRecorder) // ASSERTING the Player is a TapeRecorder
	// We know that player variable uses the interface type Player,
	// but we're pretty sure this Player is actually a TapeRecorder.
	recorder.Record()
}
