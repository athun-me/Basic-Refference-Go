package main

import (
	"context"
	"fmt"
)

// rangeChannel generates a channel sending integers from 0 to n-1.
func rangeChannel(done <-chan struct{}, n int) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

func takeFirstN(ctx context.Context, dataSource <-chan interface{}, n int) <-chan interface{} {
	takeChannel := make(chan interface{})

	go func() {
		defer close(takeChannel)
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-dataSource:
				if !ok {
					return
				}
				takeChannel <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	return takeChannel
}

func main() {
	done := make(chan struct{})
	defer close(done)

	// Generates a channel sending integers
	// From 0 to 9
	range10 := rangeChannel(done, 10)

	for num := range takeFirstN(context.Background(), range10, 5) {
		fmt.Println(num)
	}
}
