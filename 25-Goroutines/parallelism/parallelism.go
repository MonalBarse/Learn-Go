// parallelism.go
package parallelism

import (
	"fmt"
	"sync"
	"time"
)

// A function that simulates some work by sleeping for a second
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d completed\n", id)
}

// ParallelismExample demonstrates parallelism by using goroutines
func ParallelismExample() {
	var wg sync.WaitGroup

	start := time.Now()

	// Launching multiple goroutines to perform work in parallel
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("All workers completed in %v\n", duration)
}
