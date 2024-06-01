package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Statements in Golang")

	// Switch statement with a single case
	fmt.Println("Writing logic for dice of Ludo")
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // This line is used to create a source of randomness
	// here we use time.Now().UnixNano() as a seed which means the seed is the current time in nanoseconds
	// i.e the seed is always changing and hence the random number generated is always different.
	// (its like this : if the seed was same then the random number generated
	// would be same every time we run the program. Which would'nt be random at all. XD)
	// the rand.New() function creates a new Rand object that uses the provided source to generate random numbers.
	// Then rand.NewSource() creates a new pseudo-random number generator that is not safe for concurrent use by multiple goroutines.
	// The source is deterministic, so it will produce the same sequence of random numbers each time by default.
	// To produce different sequences, give it a different seed value.
	// The seed is the initial value of the generator's internal state.
	// For a given seed, the generator will produce the same sequence of numbers each time.
	// The seed value is used to initialize the generator's state.

	dice := r.Intn(6) + 1 // Generate a random number between 1 and 6
	fmt.Println("Dice rolled to:", dice)

	switch dice {
	case 1:

		fmt.Println("You got 1, You can start the game or move 1 step")
	case 2:
		fmt.Println("You rolled 2, You should move 2 steps")
	case 3:
		fmt.Println("You rolled 3, You should move 3 steps")
	case 4:
		fmt.Println("You rolled 4, You should move 4 steps")
	case 5:
		fmt.Println("You rolled 5, You should move 5 steps")
	case 6:
		fmt.Println("You rolled 6, You should move 6 steps and get another chance to roll the dice")
	default:
		fmt.Println("Invalid dice roll")
	}

}
