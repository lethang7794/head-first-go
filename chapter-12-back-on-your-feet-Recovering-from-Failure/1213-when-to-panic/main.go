package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Calling panic should be reserved for "IMPOSSIBLE" situations:
// errors that indicate a BUG in the program,
// NOT a mistake on the user's part.

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
