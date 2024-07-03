// main.go
package main

import (
	"fmt"
	"time"
)

// A function that simulates some work by sleeping for a second
func sequentialWork(id int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d completed\n", id)
}

// SequentialExample demonstrates the same task without parallelism or concurrency
func SequentialExample() {
	start := time.Now()

	// Performing work sequentially
	for i := 1; i <= 5; i++ {
		sequentialWork(i)
	}

	duration := time.Since(start)
	fmt.Printf("All workers completed sequentially in %v\n", duration)
}

func main() {
	fmt.Println("Starting Sequential Example")
	SequentialExample()

	fmt.Println("\nStarting Parallelism Example")
	ParallelismExample()

	fmt.Println("\nStarting Concurrency Example")
	ConcurrencyExample()
}
