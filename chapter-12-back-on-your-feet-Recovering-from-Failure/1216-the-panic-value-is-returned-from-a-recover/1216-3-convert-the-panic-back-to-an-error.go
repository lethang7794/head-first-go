package main

import "fmt"

func calmDown3() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown3()
	err := fmt.Errorf("there's an error")
	panic(err)
}
