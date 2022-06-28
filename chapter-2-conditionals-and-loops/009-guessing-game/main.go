// guess challenges users to guess a random number from 1 to 100
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	var success bool = false
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Your guess is HIGH.")
		} else if guess < target {
			fmt.Println("Your guess is LOW")
		} else {
			success = true
			break
		}
	}

	if success {
		fmt.Println("Good job! You did it!")
	} else {
		fmt.Println("Sorry! You didn't guess my number. It was", target)
	}
}
