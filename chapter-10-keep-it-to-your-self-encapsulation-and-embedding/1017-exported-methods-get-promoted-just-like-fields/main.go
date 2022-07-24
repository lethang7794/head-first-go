package main

import "github.com/lethang7794/head-first-go/mypackage"

func main() {
	value := mypackage.MyType{}
	value.ExportedMethod()
}
