// concurrency.go
package main

import (
	"fmt"
	"time"
)

// A function that simulates some work and sends the result to a channel
func worker(id int, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Worker %d completed", id)
}

// ConcurrencyExample demonstrates concurrency by using goroutines and channels
func ConcurrencyExample() {
	ch := make(chan string, 5)

	start := time.Now()

	// Launching multiple goroutines to perform work concurrently
	for i := 1; i <= 5; i++ {
		go worker(i, ch)
	}

	// Receiving results from the channel
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}

	duration := time.Since(start)
	fmt.Printf("All workers completed in %v\n", duration)
}
