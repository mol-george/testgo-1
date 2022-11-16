package main

import (
	"fmt"
	"math/rand"
)

// Function:
// 1. generate a random number between 1 and 100
// 2. ask the user to guess the number
// 3. if the user guesses the number, print "You win!" and the number of guesses
// 4. if the user guesses too high, print "Too high"
// 5. if the user guesses too low, print "Too low"
// 6. repeat until the user guesses the number

func guessed(secretNumber, guess int) bool {
	// use switch to check if the user guesses the number
	switch {
	case guess == secretNumber:
		fmt.Println("You win!")
		return true
	case guess > secretNumber:
		fmt.Println("Too high")
		return false
	case guess < secretNumber:
		fmt.Println("Too low")
		return false
	}
	return false
}

func main() {

	// generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	// ask the user to guess the number
	var guess int
	fmt.Println("Guess the number between 1 and 100: ")
	fmt.Scan(&guess)

	// attempt to guess the number until the user guesses the number
	for !guessed(secretNumber, guess) {
		fmt.Println("Guess again: ")
		fmt.Scan(&guess)
	}
}
