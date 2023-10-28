package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Create input and output channels.
	inputChannel := make(chan int)
	outputChannelA := make(chan int)
	outputChannelB := make(chan int)

	wg.Add(4)

	// Start the fanOut goroutine to distribute data.
	go fanOut(inputChannel, outputChannelA, outputChannelB)

	// Generate some data and send it to the input channel.
	go generateData(inputChannel)

	// Receive and print data from both output channels.

	go printData("Channel A", outputChannelA)
	go printData("Channel B", outputChannelB)

	wg.Wait()

}

func fanOut(in <-chan int, outA, outB chan int) {
	defer wg.Done()
	for data := range in {
		select {
		case outA <- data:
			// Data is sent to outA.
		case outB <- data:
			// Data is sent	 to outB.
		}
	}

	// Close the output channels when the input channel is closed.
	close(outA)
	close(outB)
}

func generateData(ch chan<- int) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // Close the input channel when done.
}

func printData(label string, ch <-chan int) {
	defer wg.Done()

	for data := range ch {
		fmt.Printf("%s received: %d\n", label, data)
	}
}
