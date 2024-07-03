// mutex/mutex.go
package mutex

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int // Shared counter
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

// Increment increments the counter using mutex for synchronization
func Increment(id int) {
	defer wg.Done()

	mutex.Lock()
	counter++
	fmt.Printf("Incremented counter by worker %d to %d\n", id, counter)
	mutex.Unlock()
}

// MutexExample demonstrates using mutex for synchronization
func MutexExample() {
	start := time.Now()

	// Number of goroutines to launch
	numWorkers := 5

	// Add goroutines to wait group
	wg.Add(numWorkers)

	// Launch multiple goroutines to increment the counter
	for i := 1; i <= numWorkers; i++ {
		go Increment(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Final counter expected value 5; actual value: %d\n", counter)
	fmt.Printf("All workers completed in %v\n", duration)
}

// Without Mutex this exapmle will not work as expected
// Here is an example of the same program without mutex

/*
package  mutex

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int // Shared counter
	wg      sync.WaitGroup
)

func worker(id int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work
	time.Sleep(100 * time.Millisecond)

	// Increment the counter
	counter++
	fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
}

func MutexExample() {
	// Number of goroutines to launch
	numWorkers := 5

	// Add goroutines to wait group
	wg.Add(numWorkers)

	// Launch multiple workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	// Wait for all workers to finish
	wg.Wait()
		duration := time.Since(start)

	fmt.Printf("Final counter value: %d\n", counter)
}

*/