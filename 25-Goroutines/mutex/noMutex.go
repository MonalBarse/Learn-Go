// mutex/no_mutex_example.go

package mutex

import (
	"fmt"
	"sync"
)

var (
	counterNoMutex int // Shared counter without mutex
	wgNoMutex      sync.WaitGroup
)

// IncrementNoMutex increments the counter without using mutex
func IncrementNoMutex(id int) {
	defer wgNoMutex.Done()

	// Simulate some work
	// Note: Removed time.Sleep to allow concurrent execution
	// time.Sleep(100 * time.Millisecond)

	// Access and increment shared counter without synchronization
	counterNoMutex++
	fmt.Printf("Worker %d incremented counter (without mutex) to %d\n", id, counterNoMutex)
}

// NoMutexExample demonstrates the race condition without using mutex
func NoMutexExample() {
	// Number of goroutines to launch
	numWorkers := 5

	// Add goroutines to wait group
	wgNoMutex.Add(numWorkers)

	// Launch multiple goroutines to increment the counter without mutex
	for i := 1; i <= numWorkers; i++ {
		go IncrementNoMutex(i)
	}

	// Wait for all goroutines to finish
	wgNoMutex.Wait()
	fmt.Printf("Final counter expected value 5; actual value: %d (without mutex)\n", counterNoMutex)
	fmt.Println("All workers completed (without mutex)")
	fmt.Printf("Run ` go run -race main.go ` to execute main.go with the race detector enabled\n")

}
