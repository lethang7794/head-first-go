package main

import "fmt"

func main() {
	const num = 12.3456
	fmt.Printf("%%7.3f: %7.3f\n", num)
	fmt.Printf("%%7.2f: %7.2f\n", num)
	fmt.Printf("%%7.1f: %7.1f\n", num)

	fmt.Printf(" %%.1f: %.1f\n", num)
	fmt.Printf(" %%.2f: %.2f\n", num)
	fmt.Printf(" %%.3f: %.3f\n", num)
	fmt.Printf(" %%.4f: %.4f\n", num)
	fmt.Printf(" %%.5f: %.5f\n", num)
	fmt.Printf(" %%.6f: %.6f\n", num)

	fmt.Printf("%.2f\n", 1.26000000000000000002)
	fmt.Printf("%.2f\n", 1.81999999999999999998)

}
