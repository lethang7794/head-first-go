package main

import "github.com/lethang7794/head-first-go/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("A Test Track")
	player.Stop()

	// Assign the second return value to a variable
	recorder, ok := player.(gadget.TapeRecorder)
	// Call the Record method only if the original value was a TapeRecorder
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(gadget.TapeRecorder{}) // TryOut function works with a TapeRecorder value
	TryOut(gadget.TapePlayer{})   // and a TapePlayer value too.
}
