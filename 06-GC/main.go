// Here is my Blog on GC: http://link.medium.com/W85XK1OU2Jb
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// Set GOGC to 200
	os.Setenv("GOGC", "200")
	/*
	   os.Setenv("GOGC", "200"): This sets the environment variable GOGC to 200.
	   The GOGC environment variable controls the aggressiveness of the garbage
	   collector. Setting GOGC to 200 means the heap can grow by 200% before
	   triggering a garbage collection cycle. (The default value is 100)
	*/

	// Create a large number of objects
	for i := 0; i < 1000000; i++ {
		_ = make([]int, 0, 1000)
	}

	// Print memory statistics
	var m runtime.MemStats //struct that holds various memory statistics.
	runtime.ReadMemStats(&m)
	// m.Alloc is the number of bytes allocated and not yet freed.
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))           //Prints the currently allocated heap objects in MB
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc)) // Prints the total bytes allocated for heap object
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))               // Prints the total bytes of memory obtained from the OS.
	fmt.Printf("NumGC = %v\n", m.NumGC)                      // Prints the number of completed garbage collection cycles.
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024 // Converts bytes to MB
}
