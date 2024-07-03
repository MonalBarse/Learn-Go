// main.go

package main

import (
	"fmt"
	"time"

	"github.com/MonalBarse/Learn-Go/tree/main/25-Goroutines/channels"
	"github.com/MonalBarse/Learn-Go/tree/main/25-Goroutines/concurrency"
	"github.com/MonalBarse/Learn-Go/tree/main/25-Goroutines/mutex"
	"github.com/MonalBarse/Learn-Go/tree/main/25-Goroutines/parallelism"
)

func main() {
	fmt.Println("Learning about Goroutines, Parallelism, Concurrency, Mutex, and Channels in Go")

	fmt.Println("\n--- Parallelism Example ---")

	parallelism.ParallelismExample()

	time.Sleep(500 * time.Millisecond) // Add delay between examples

	fmt.Println("\n--- Concurrency Example ---")
	concurrency.ConcurrencyExample()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--- Mutex Example ---")
	mutex.MutexExample() // Use MutexExample from the mutex package

	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--- No Mutex Example (Race Condition) ---")
	mutex.NoMutexExample() // Use NoMutexExample from the mutex package

	// --- Channels Example ---
	channels.RunChannelsExample()
	fmt.Println("\nCompleted learning examples")
}
