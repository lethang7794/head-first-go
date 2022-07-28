package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) { // Has a Play method just like TapePlayer's
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() { // Has a Stop method just like TapePlayer's
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
