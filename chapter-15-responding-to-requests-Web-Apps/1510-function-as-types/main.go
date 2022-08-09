package main

import (
	"fmt"
	"net/http"
)

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiple(a int, b int) float64 {
	return float64(a) * float64(b)
}

func doMath(fn func(int, int) float64) {
	a, b := 10, 20
	result := fn(a, b)
	fmt.Printf("fn(%d, %d) = %0.2f\n", a, b, result)
}

func main() {
	var greeterFunction func()
	var mathFunction func(int, int) float64

	greeterFunction = sayHi
	mathFunction = divide

	//greeterFunction = divide // Cannot use 'divide' (type func(a int, b int) float64) as the type func()
	//mathFunction = sayHi // Cannot use 'sayHi' (type func()) as the type func(int, int) float64

	greeterFunction()
	fmt.Println(mathFunction(1, 10))

	doMath(divide)
	doMath(multiple)

	//http.HandleFunc("/", sayHi) // Cannot use 'sayHi' (type func()) as the type func(ResponseWriter, *Request)
	http.HandleFunc("/", handler)
}

func handler(writer http.ResponseWriter, request *http.Request) {

}
