package channels

import (
	"fmt"
	"sync"
)

// channels, synchronization, and potential deadlocks
func RunChannelsExample() {
	fmt.Println("------ Enhanced Channels in Golang ------")

	ch := make(chan int, 2) // Create a buffered channel with a capacity of 2
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go sendValues(ch, wg, 1, 2, 3, 4)
	go sendValues(ch, wg, 5, 6, 7, 8)
	go receiveValues(ch, wg)
	go receiveValues(ch, wg)

	wg.Wait()

	// Demonstrate closing a channel and handling closure
	close(ch)
	fmt.Println("Channel closed, no more values can be sent.")

	fmt.Println("Trying to receive from a closed channel")
	wg.Add(1)
	go receiveAfterClose(ch, wg)
	wg.Wait()
}

func sendValues(ch chan<- int, wg *sync.WaitGroup, values ...int) {
	defer wg.Done()
	for _, v := range values {
		ch <- v
		fmt.Printf("Sent %d to channel\n", v)
	}
}

func receiveValues(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ { // Receive 4 values
		v := <-ch
		fmt.Printf("Received %d from channel\n", v)
	}
}

func receiveAfterClose(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		value, isChannelOpen := <-ch
		fmt.Printf("Is Channel Open: %v\n", isChannelOpen)
		if !isChannelOpen {
			fmt.Println("Channel closed, no more values to receive.")
			break
		}
		fmt.Printf("Received %d from closed channel\n", value)
	}
}
