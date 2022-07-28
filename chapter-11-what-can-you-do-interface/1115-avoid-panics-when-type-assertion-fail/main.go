package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func main() {
	var player Player = gadget.TapePlayer{}
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}
