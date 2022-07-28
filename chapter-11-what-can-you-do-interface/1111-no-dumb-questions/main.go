package main

// ? Should interface type names begin with a capital letter or a lowercase letter ?

type unexportedInterface interface {
	DoSomething()
}
type ExportedInterface interface {
	DoAnotherThing()
}

func aFunction(u unexportedInterface) {
	u.DoSomething()
}

func AnOtherFunction(e ExportedInterface) {
	e.DoAnotherThing()
}

type AType string

func (a AType) DoSomething() {}

type AnotherType int

func (a AnotherType) DoAnotherThing() {}

func main() {
	var u unexportedInterface
	u = AType("AType")
	aFunction(u)

	var e ExportedInterface
	e = AnotherType(1)
	AnOtherFunction(e)
}
