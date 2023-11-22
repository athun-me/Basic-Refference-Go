package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	queue := make(chan bool, 3)
	for _, id := range arr {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			queue <- true
			defer func() { <-queue }()
			makeRequest(id)
		}(id)
	}

	fmt.Println("number of goroutine", runtime.NumGoroutine())
	wg.Wait()
}

func makeRequest(id int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("id :", id)
}
